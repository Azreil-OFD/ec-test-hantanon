package profile

// import (
// 	"backend/internal/database"
// 	"backend/internal/middleware"
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	// "github.com/jackc/pgx"
// )

// // GetFriendsInfoHandler - универсальная ручка для получения информации о друзьях и заявках
// func GetFriendsInfoHandler(w http.ResponseWriter, r *http.Request) {
// 	// Получаем UUID текущего пользователя из контекста
// 	userID := r.Context().Value(middleware.UserUUIDKey).(string)

// 	// Получаем тип запроса из параметров
// 	requestType := r.URL.Query().Get("type")

// 	// Если параметр type не передан, будем возвращать все категории
// 	if requestType == "" {
// 		// Собираем все результаты
// 		friends, err := getFriends(userID)
// 		if err != nil {
// 			http.Error(w, "Ошибка при получении списка друзей", http.StatusInternalServerError)
// 			return
// 		}

// 		incoming, err := getIncomingRequests(userID)
// 		if err != nil {
// 			http.Error(w, "Ошибка при получении входящих заявок", http.StatusInternalServerError)
// 			return
// 		}

// 		outgoing, err := getOutgoingRequests(userID)
// 		if err != nil {
// 			http.Error(w, "Ошибка при получении исходящих заявок", http.StatusInternalServerError)
// 			return
// 		}

// 		// Возвращаем все в одном ответе
// 		result := map[string][]UserProfile{
// 			"friends":  friends,
// 			"incoming": incoming,
// 			"outgoing": outgoing,
// 		}

// 		// Отправляем результат в формате JSON
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(result)

// 		return
// 	}

// 	// Если параметр type передан, обрабатываем его как раньше
// 	var query string
// 	var rows *pgx.Rows
// 	var err error
// 	result := []UserProfile{}

// 	switch requestType {
// 	case "friends":
// 		query = `SELECT u.id, u.login, u.full_name, u.email
// 				  FROM users u
// 				  JOIN friends f ON (u.id = f.friend_id AND f.user_id = $1 AND f.status = 'accepted') 
// 				  OR (u.id = f.user_id AND f.friend_id = $1 AND f.status = 'accepted')`
// 		rows, err = database.DB.Query(context.Background(), query, userID)
// 		if err != nil {
// 			log.Println("Ошибка при получении списка друзей:", err)
// 			http.Error(w, "Ошибка при получении списка друзей", http.StatusInternalServerError)
// 			return
// 		}
// 		defer rows.Close()

// 		for rows.Next() {
// 			var friend UserProfile
// 			if err := rows.Scan(&friend.UUID, &friend.Login, &friend.FullName, &friend.Email); err != nil {
// 				log.Println("Ошибка при сканировании данных:", err)
// 				http.Error(w, "Ошибка при сканировании данных", http.StatusInternalServerError)
// 				return
// 			}
// 			result = append(result, friend)
// 		}

// 	case "incoming":
// 		query = `SELECT u.id, u.login, u.full_name, u.email
// 				  FROM users u
// 				  JOIN friends f ON u.id = f.user_id
// 				  WHERE f.friend_id = $1 AND f.status = 'pending'`
// 		rows, err = database.DB.Query(context.Background(), query, userID)
// 		if err != nil {
// 			log.Println("Ошибка при получении входящих заявок:", err)
// 			http.Error(w, "Ошибка при получении входящих заявок", http.StatusInternalServerError)
// 			return
// 		}
// 		defer rows.Close()

// 		for rows.Next() {
// 			var user UserProfile
// 			if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
// 				log.Println("Ошибка при сканировании данных:", err)
// 				http.Error(w, "Ошибка при сканировании данных", http.StatusInternalServerError)
// 				return
// 			}
// 			result = append(result, user)
// 		}

// 	case "outgoing":
// 		query = `SELECT u.id, u.login, u.full_name, u.email
// 				  FROM users u
// 				  JOIN friends f ON u.id = f.friend_id
// 				  WHERE f.user_id = $1 AND f.status = 'pending'`
// 		rows, err = database.DB.Query(context.Background(), query, userID)
// 		if err != nil {
// 			log.Println("Ошибка при получении исходящих заявок:", err)
// 			http.Error(w, "Ошибка при получении исходящих заявок", http.StatusInternalServerError)
// 			return
// 		}
// 		defer rows.Close()

// 		for rows.Next() {
// 			var user UserProfile
// 			if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
// 				log.Println("Ошибка при сканировании данных:", err)
// 				http.Error(w, "Ошибка при сканировании данных", http.StatusInternalServerError)
// 				return
// 			}
// 			result = append(result, user)
// 		}

// 	default:
// 		http.Error(w, "Некорректный параметр 'type'. Ожидается 'friends', 'incoming' или 'outgoing'", http.StatusBadRequest)
// 		return
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Println("Ошибка при обработке строк:", err)
// 		http.Error(w, "Ошибка при обработке строк", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(result)
// }

// // Функция для получения списка друзей
// func getFriends(userID string) ([]UserProfile, error) {
// 	query := `SELECT u.id, u.login, u.full_name, u.email
// 			  FROM users u
// 			  JOIN friends f ON (u.id = f.friend_id AND f.user_id = $1 AND f.status = 'accepted') 
// 			  OR (u.id = f.user_id AND f.friend_id = $1 AND f.status = 'accepted')`
// 	rows, err := database.DB.Query(context.Background(), query, userID)
// 	if err != nil {
// 		log.Println("Ошибка при получении списка друзей:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	users := []UserProfile{}
// 	for rows.Next() {
// 		var user UserProfile
// 		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
// 			log.Println("Ошибка при сканировании данных:", err)
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Println("Ошибка при обработке строк:", err)
// 		return nil, err
// 	}

// 	return users, nil
// }

// // Функция для получения списка входящих заявок
// func getIncomingRequests(userID string) ([]UserProfile, error) {
// 	query := `SELECT u.id, u.login, u.full_name, u.email
// 			  FROM users u
// 			  JOIN friends f ON u.id = f.user_id
// 			  WHERE f.friend_id = $1 AND f.status = 'pending'`
// 	rows, err := database.DB.Query(context.Background(), query, userID)
// 	if err != nil {
// 		log.Println("Ошибка при получении входящих заявок:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	users := []UserProfile{}
// 	for rows.Next() {
// 		var user UserProfile
// 		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
// 			log.Println("Ошибка при сканировании данных:", err)
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Println("Ошибка при обработке строк:", err)
// 		return nil, err
// 	}

// 	return users, nil
// }

// // Функция для получения списка исходящих заявок
// func getOutgoingRequests(userID string) ([]UserProfile, error) {
// 	query := `SELECT u.id, u.login, u.full_name, u.email
// 			  FROM users u
// 			  JOIN friends f ON u.id = f.friend_id
// 			  WHERE f.user_id = $1 AND f.status = 'pending'`
// 	rows, err := database.DB.Query(context.Background(), query, userID)
// 	if err != nil {
// 		log.Println("Ошибка при получении исходящих заявок:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	users := []UserProfile{}
// 	for rows.Next() {
// 		var user UserProfile
// 		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
// 			log.Println("Ошибка при сканировании данных:", err)
// 			return nil, err
// 		}
// 		users = append(users, user)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Println("Ошибка при обработке строк:", err)
// 		return nil, err
// 	}

// 	return users, nil
// }
