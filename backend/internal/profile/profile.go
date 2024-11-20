package profile

import (
	"backend/internal/middleware"
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

// UserProfile представляет профиль пользователя
// @Description Структура для профиля пользователя
type UserProfile struct {
	UUID     string `json:"uuid"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

// GetProfile godoc
// @Summary Получение профиля пользователя по UUID
// @Description Возвращает профиль пользователя на основе переданного UUID. Если пользователь не найден, возвращается ошибка 404.
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param uuid path string true "UUID пользователя"  // Здесь мы ожидаем UUID в пути запроса
// @Success 200 {object} UserProfile "Профиль пользователя"  // Возвращаем профиль пользователя в ответе
// @Failure 404 {string} string "Пользователь не найден"  // Ошибка, если пользователь не найден
// @Failure 500 {string} string "Внутренняя ошибка сервера"  // Ошибка сервера
// @Router /api/profile/{uuid} [get]
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
