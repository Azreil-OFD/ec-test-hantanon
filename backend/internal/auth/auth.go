package auth

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/internal/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type request struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var request request
	response := model.Response{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Message = "Неверное тело запроса"
		response.Status = http.StatusBadRequest
		model.SendJSONResponse(w, response)
		return
	}

	user, err := service.GetUserByLogin(request.Login)
	if err != nil {
		if customErr, ok := err.(*model.CustomError); ok {
			if customErr.Code == model.NotFound {
				response.Message = "Неверный логин или пароль"
				response.Status = http.StatusNotFound
			} else if customErr.Code == model.DBError {
				response.Message = customErr.Message
				response.Status = http.StatusInternalServerError
			}
		}
		model.SendJSONResponse(w, response)
		return
	}

	if !util.ComparePassword(user.PasswordHash, request.Password) {
		response.Message = "Неверный логин или пароль"
		response.Status = http.StatusNotFound
		model.SendJSONResponse(w, response)
		return
	}
	token, err := util.GenerateJWT(user.UUID)
	if err != nil {
		response.Message = "Ошибка генерации токена"
		response.Status = http.StatusInternalServerError
		model.SendJSONResponse(w, response)
		return
	}

	response.Message = fmt.Sprintf("Пользователь %s успешно авторизовался!", user.Login)
	response.Status = http.StatusOK
	response.Data = map[string]string{
		"token": token,
	}
	model.SendJSONResponse(w, response)
}


func TestHandler(w http.ResponseWriter, r *http.Request) {
	model.SendJSONResponse(w, model.Response{
		Message: "тест",
		Status: 200,
	})
}