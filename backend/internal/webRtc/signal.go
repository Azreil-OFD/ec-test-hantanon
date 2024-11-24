package webrtc

import (
	"backend/internal/middleware"
	"backend/internal/model"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Клиентский тип
type Client struct {
	conn *websocket.Conn
	send chan []byte
}

// Хранение клиентов, каждый клиент получает уникальный идентификатор
var clients = make(map[string]*Client)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type request struct {
	UUIDs []string    `json:"uuid"`
	Type  string      `json:"type"`
	SDP   interface{} `json:"sdp"`
}

// Обработчик для сигнального канала WebSocket
func SignalHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем userID из контекста
	userID, _ := r.Context().Value(middleware.UserUUIDKey).(string)

	// Устанавливаем WebSocket-соединение
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка подключения WebSocket:", err)
		http.Error(w, "Не удалось установить WebSocket соединение", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Создание нового клиента и добавление его в список
	client := &Client{conn: conn, send: make(chan []byte)}
	clients[userID] = client

	// Обработка входящих сообщений от клиента
	go handleMessages(client)

	// Чтение сообщений от клиента
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			break
		}

		// Разбор сообщения
		var req request
		if err := json.NewDecoder(bytes.NewReader(msg)).Decode(&req); err != nil {
			model.SendJSONResponse(w, model.Response{
				Message: "Неверное тело запроса",
				Status:  http.StatusBadRequest,
			})
			continue
		}
		for _, uuid := range req.UUIDs {
			if otherClient, ok := clients[uuid]; ok {
				response := struct {
					UUID string      `json:"uuid"`
					Type string      `json:"type"`
					SDP  interface{} `json:"sdp"`
				} {
					UUID: userID,
					Type: req.Type,
					SDP: req.SDP,
				}
				responseMsg, err := json.Marshal(response)
				if err != nil {
					log.Println("Ошибка сериализации структуры в JSON:", err)
					continue // если не удается сериализовать, продолжаем ждать следующее сообщение
				}
				err = otherClient.conn.WriteMessage(websocket.TextMessage, responseMsg)
				if err != nil {
					log.Printf("Ошибка отправки сообщения клиенту %s: %v", uuid, err)
				}
			} else {
				model.SendJSONResponse(w, model.Response{
					Message: "",
				})
				log.Printf("Клиент с UUID %s не найден", uuid)
			}
		}
	}
	delete(clients, userID)
}

// Обработка сообщений (например, передача ICE-кандидатов)
func handleMessages(client *Client) {
	for {
		select {
		case msg := <-client.send:

			err := client.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Ошибка отправки сообщения:", err)
			}
		}
	}
}
