package registration

import (
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

type request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	response := model.Response{}
	if r.Method != http.MethodPost {
		response.Message = "Неверный метод запроса"
		response.Status = http.StatusMethodNotAllowed
		model.SendJSONResponse(w, response)
		return
	}

	var request request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = "Неверное тело запроса"
		response.Status = http.StatusBadRequest
		model.SendJSONResponse(w, response)
		return
	}

	if strings.TrimSpace(request.Login) == "" || strings.TrimSpace(request.Password) == "" || strings.TrimSpace(request.Email) == "" || strings.TrimSpace(request.FullName) == "" {
		response.Message = "Логин, пароль, email и имя обязательны"
		response.Status = http.StatusBadRequest
		model.SendJSONResponse(w, response)
		return
	}

	request.Password, err = util.HashPassword(request.Password)
	if err != nil {
		response.Message = "Ошибка при хешировании пароля"
		response.Status = http.StatusInternalServerError
		model.SendJSONResponse(w, response)
		return
	}

	err = saveUserToDB(request)
	if err != nil {
		if customErr, ok := err.(*model.CustomError); ok {
			if customErr.Code == model.ConflictError {
				response.Status = http.StatusConflict
			} else if customErr.Code == model.DBError {
				response.Status = http.StatusInternalServerError
			}
			response.Message = customErr.Message
		}
		model.SendJSONResponse(w, response)
		return
	}

	response.Message = fmt.Sprintf("Пользователь %s успешно зарегистрирован!", request.Login)
	response.Status = http.StatusCreated
	model.SendJSONResponse(w, response)
}

func saveUserToDB(user request) error {
	query := `INSERT INTO users (login, password, email, full_name) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(context.Background(), query, user.Login, user.Password, user.Email, user.FullName)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				if strings.Contains(pgErr.Message, "login") {
					err = &model.CustomError{
						Message: "Логин уже существует",
						Code:    model.ConflictError,
					}
				} else if strings.Contains(pgErr.Message, "email") {
					err = &model.CustomError{
						Message: "Email уже существует",
						Code:    model.ConflictError,
					}
				} else {
					log.Printf("ВАЖНОЕ!!! ошибка: %v\nLogin: %s\nPassword: %s\nEmail: %s\nFullName: %s", err, user.Login, user.Password, user.Email, user.FullName)
					err = &model.CustomError{
						Message: "Кто ты, воин?",
						Code:    model.ConflictError,
					}
				}
				return err
			}
		}
		log.Println("Ошибка при вставке пользователя в базу данных:", err)
		return &model.CustomError{
			Message: "Ошибка на стороне сервера",
			Code:    model.DBError,
		}
	}
	return nil
}
