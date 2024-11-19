package profile

import (
	"backend/internal/database"
	"backend/internal/middleware"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GetFriendsInfoHandler - универсальная ручка для получения информации о друзьях и заявках
func GetFriendsInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем UUID текущего пользователя из контекста
	userID := r.Context().Value(middleware.UserUUIDKey).(string)

	// Получаем тип запроса из параметров
	requestType := r.URL.Query().Get("type")

	// Если параметр type не передан, возвращаем все категории
	if requestType == "" {
		// Собираем все результаты
		friends, err := getUserListByType(userID, "friends")
		if err != nil {
			http.Error(w, "Ошибка при получении списка друзей", http.StatusInternalServerError)
			return
		}

		incoming, err := getUserListByType(userID, "incoming")
		if err != nil {
			http.Error(w, "Ошибка при получении входящих заявок", http.StatusInternalServerError)
			return
		}

		outgoing, err := getUserListByType(userID, "outgoing")
		if err != nil {
			http.Error(w, "Ошибка при получении исходящих заявок", http.StatusInternalServerError)
			return
		}

		// Возвращаем все в одном ответе
		result := map[string][]UserProfile{
			"friends":  friends,
			"incoming": incoming,
			"outgoing": outgoing,
		}

		// Отправляем результат в формате JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

		return
	}

	// Если параметр type передан, обрабатываем его с помощью функции
	result, err := getUserListByType(userID, requestType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем результат в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getUserListByType(userID, requestType string) ([]UserProfile, error) {
	switch requestType {
	case "friends":
		return getFriends(userID)
	case "incoming":
		return getIncomingRequests(userID)
	case "outgoing":
		return getOutgoingRequests(userID)
	default:
		return nil, fmt.Errorf("Некорректный параметр 'type'. Ожидается 'friends', 'incoming' или 'outgoing'")
	}
}

// Функция для получения списка друзей
func getFriends(userID string) ([]UserProfile, error) {
	query := `SELECT u.id, u.login, u.full_name, u.email
			  FROM users u
			  JOIN friends f ON (u.id = f.friend_id AND f.user_id = $1 AND f.status = 'accepted') 
			  OR (u.id = f.user_id AND f.friend_id = $1 AND f.status = 'accepted')`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Println("Ошибка при получении списка друзей:", err)
		return nil, err
	}
	defer rows.Close()

	users := []UserProfile{}
	for rows.Next() {
		var user UserProfile
		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
			log.Println("Ошибка при сканировании данных:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Ошибка при обработке строк:", err)
		return nil, err
	}

	return users, nil
}

// Функция для получения списка входящих заявок
func getIncomingRequests(userID string) ([]UserProfile, error) {
	query := `SELECT u.id, u.login, u.full_name, u.email
			  FROM users u
			  JOIN friends f ON u.id = f.user_id
			  WHERE f.friend_id = $1 AND f.status = 'pending'`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Println("Ошибка при получении входящих заявок:", err)
		return nil, err
	}
	defer rows.Close()

	users := []UserProfile{}
	for rows.Next() {
		var user UserProfile
		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
			log.Println("Ошибка при сканировании данных:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Ошибка при обработке строк:", err)
		return nil, err
	}

	return users, nil
}

// Функция для получения списка исходящих заявок
func getOutgoingRequests(userID string) ([]UserProfile, error) {
	query := `SELECT u.id, u.login, u.full_name, u.email
			  FROM users u
			  JOIN friends f ON u.id = f.friend_id
			  WHERE f.user_id = $1 AND f.status = 'pending'`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Println("Ошибка при получении исходящих заявок:", err)
		return nil, err
	}
	defer rows.Close()

	users := []UserProfile{}
	for rows.Next() {
		var user UserProfile
		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
			log.Println("Ошибка при сканировании данных:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Ошибка при обработке строк:", err)
		return nil, err
	}

	return users, nil
}
