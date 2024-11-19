package profile

import (
	"backend/internal/middleware"
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

// Структура для профиля пользователя
type UserProfile struct {
	UUID     string `json:"uuid"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

// Функция для получения профиля пользователя по UUID
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// Извлекаем UUID пользователя из контекста запроса
	uuid, ok := r.Context().Value(middleware.UserUUIDKey).(string)
	if !ok || uuid == "" {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	// Получаем данные пользователя по UUID из базы данных
	user, err := service.GetUserByUUID(uuid)
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	// Отправляем профиль пользователя в ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
