package profile

import (
	"backend/internal/database"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Функция для поиска пользователей с пагинацией
func SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр поиска из запроса
	searchQuery := r.URL.Query().Get("query")
	if searchQuery == "" {
		http.Error(w, "Параметр запроса 'query' обязателен", http.StatusBadRequest)
		return
	}

	// Определяем параметры пагинации
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		page = 1 // По умолчанию первая страница
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10 // По умолчанию 10 результатов на странице
	}

	users, err := searchUsers(searchQuery, page, limit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем найденных пользователей в ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// Функция для поиска пользователей по имени (часть имени)
func searchUsers(query string, page, limit int) ([]UserProfile, error) {
	offset := (page - 1) * limit
	// Выполняем поиск как по логину, так и по полному имени
	queryString := `SELECT id, login, full_name, email 
					FROM users 
					WHERE login ILIKE $1 OR full_name ILIKE $2 
					ORDER BY full_name, login 
					LIMIT $3 OFFSET $4`
	rows, err := database.DB.Query(context.Background(), queryString, "%"+query+"%", "%"+query+"%", limit, offset)

	if err != nil {
		log.Println("Ошибка при поиске пользователя:", err)
		return nil, err
	}
	defer rows.Close()

	// Обрабатываем результаты запроса
	users := []UserProfile{}
	for rows.Next() {
		var user UserProfile
		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
			log.Println("Ошибка при сканировании данных:", err)
			return nil, err
		}
		users = append(users, user)
	}

	// Обрабатываем возможные ошибки после перебора строк
	if err := rows.Err(); err != nil {
		log.Println("Ошибка при обработке строк:", err)
		return nil, err
	}

	return users, nil
}
