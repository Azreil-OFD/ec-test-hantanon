package profile

import (
	"backend/internal/database"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// SearchUserHandler godoc
// @Summary Поиск пользователей
// @Description Выполняет поиск пользователей по переданному запросу. Поддерживает пагинацию с параметрами "page" и "limit".
// @Tags Users
// @Accept json
// @Produce json
// @Param query query string true "Строка для поиска"  // Параметр поиска обязательный
// @Param page query int false "Номер страницы" default(1) // Параметр пагинации, по умолчанию 1
// @Param limit query int false "Количество пользователей на странице" default(10) // Параметр ограничения по количеству результатов на странице
// @Success 200 {array} UserProfile "Список пользователей"  // Возвращаем список пользователей
// @Failure 400 {string} string "Ошибка: Параметр 'query' обязателен"  // Ошибка, если параметр поиска не передан
// @Failure 500 {string} string "Внутренняя ошибка сервера"  // Ошибка сервера
// @Router /api/search [get]
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
