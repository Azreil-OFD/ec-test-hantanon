package profile

import (
	"backend/internal/middleware"
	"backend/internal/model"
	"backend/internal/service"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	response := model.Response{}

	uuid, _ := r.Context().Value(middleware.UserUUIDKey).(string)

	user, err := service.GetUserByUUID(uuid)
	if err != nil {
		customErr, _ := err.(*model.CustomError)
		response.Status = customErr.Code
		response.Message = customErr.Message
		model.SendJSONResponse(w, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Запрос на получение профиля прошел успешно!"
	response.Data = user
	model.SendJSONResponse(w, response)
}
