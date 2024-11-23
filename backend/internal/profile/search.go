package profile

import (
	"backend/internal/database"
	"backend/internal/model"
	"context"
	"log"
	"net/http"
	"strconv"
)

func SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	response := model.Response{}
	searchQuery := r.URL.Query().Get("query")
	if searchQuery == "" {
		response.Message = "Параметр запроса 'query' обязателен"
		response.Status = model.BadRequest
		model.SendJSONResponse(w, response)
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

	users, totalPages, err := searchUsers(searchQuery, page, limit)
	if err != nil {
		customErr, _ := err.(*model.CustomError)
		response.Status = customErr.Code
		response.Message = customErr.Message
		model.SendJSONResponse(w, response)
		return
	}

	response.Message = "Запрос на поиск прошел успешно!"
	response.Status = http.StatusOK
	response.Data = struct {
		MetaData struct {
			CurentLimit int `json:"curent_limit"`
			TotalPages  int `json:"total_pages"`
		} `json:"metadata"`
		Data interface{} `json:"data"`
	}{
		MetaData: struct {
			CurentLimit int `json:"curent_limit"`
			TotalPages  int `json:"total_pages"`
		}{
			CurentLimit: len(users),
			TotalPages:  totalPages,
		},
		Data: users,
	}
	model.SendJSONResponse(w, response)
}

// Функция для поиска пользователей по имени (часть имени)
func searchUsers(query string, page, limit int) ([]model.User, int, error) {
	// Подсчитываем общее количество пользователей, удовлетворяющих запросу
	var totalCount int
	countQuery := `SELECT COUNT(*) 
				   FROM users 
				   WHERE login ILIKE $1 OR full_name ILIKE $2`
	err := database.DB.QueryRow(context.Background(), countQuery, "%"+query+"%", "%"+query+"%").Scan(&totalCount)
	if err != nil {
		log.Println("Ошибка: ", err)
		return nil, 0, &model.CustomError{
			Code:    model.DBError,
			Message: "Ошибка на стороне сервера",
		}
	}

	offset := (page - 1) * limit
	// Выполняем поиск как по логину, так и по полному имени
	queryString := `SELECT id, login, full_name, email 
					FROM users 
					WHERE login ILIKE $1 OR full_name ILIKE $2 
					ORDER BY full_name, login 
					LIMIT $3 OFFSET $4`
	rows, err := database.DB.Query(context.Background(), queryString, "%"+query+"%", "%"+query+"%", limit, offset)
	if err != nil {
		log.Println("Ошибка: ", err)
		return nil, 0, &model.CustomError{
			Code:    model.DBError,
			Message: "Ошибка на стороне сервера",
		}
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
			log.Println("Ошибка: ", err)
			return nil, 0, &model.CustomError{
				Code:    model.DBError,
				Message: "Ошибка на стороне сервера",
			}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Ошибка: ", err)
		return nil, 0, &model.CustomError{
			Code:    model.DBError,
			Message: "Ошибка на стороне сервера",
		}
	}

	// Рассчитываем общее количество страниц
	totalPages := (totalCount + limit - 1) / limit // Это дает количество страниц с округлением в большую сторону

	return users, totalPages, nil
}
