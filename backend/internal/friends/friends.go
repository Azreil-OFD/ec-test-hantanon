package friends

import (
	"backend/internal/middleware"
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

// SendFriendRequestHandler godoc
// @Summary Отправить запрос на добавление в друзья
// @Description Отправляет запрос на добавление в друзья пользователю с указанным логином. Требуется JWT токен в заголовке.
// @Tags Friends
// @Accept json
// @Produce json
// @Param friend_login query string true "Логин друга" 
// @Success 200 {object} map[string]string {"message": "Регистрация друга успешна"}
// @Failure 400 {string} string "Логин друга обязателен"
// @Failure 404 {string} string "Друг не найден"
// @Failure 500 {string} string "Ошибка при добавлении друга"
// @Security BearerAuth
// @Router /api/friends/request [post]
func SendFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем UUID текущего пользователя из контекста
	userID := r.Context().Value(middleware.UserUUIDKey).(string)

	// Получаем friendID из параметров запроса
	friendLogin := r.URL.Query().Get("friend_login")
	if friendLogin == "" {
		http.Error(w, "Логин друга обязателен", http.StatusBadRequest)
		return
	}
	friendUser, err := service.GetUserByLogin(friendLogin)
	if err != nil {
		http.Error(w, "Друг не найден", http.StatusNotFound)
	}
	err = sendFriendRequest(userID, friendUser.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Регистрация друга успешна"})
}

// AcceptFriendRequestHandler godoc
// @Summary Принять запрос на добавление в друзья
// @Description Принять запрос на добавление в друзья от пользователя с указанным логином. Требуется JWT токен в заголовке.
// @Tags Friends
// @Accept json
// @Produce json
// @Param friend_login query string true "Логин друга"
// @Success 200 {object} map[string]string {"message": "Запрос на добавление в друзья принят"}
// @Failure 400 {string} string "Логин друга обязателен"
// @Failure 404 {string} string "Друг не найден"
// @Failure 500 {string} string "Ошибка при принятии запроса"
// @Security BearerAuth
// @Router /api/friends/accept [post]
func AcceptFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем UUID текущего пользователя из контекста
	userID := r.Context().Value(middleware.UserUUIDKey).(string)

	// Получаем friendID из параметров запроса
	friendLogin := r.URL.Query().Get("friend_login")
	if friendLogin == "" {
		http.Error(w, "Логин друга обязателен", http.StatusBadRequest)
		return
	}

	// Получаем информацию о друге
	friendUser, err := service.GetUserByLogin(friendLogin)
	if err != nil {
		http.Error(w, "Друг не найден", http.StatusNotFound)
		return
	}

	// Обновляем статус запроса на принятие
	err = acceptFriendRequest(userID, friendUser.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ответ клиенту
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Запрос на добавление в друзья принят"})
}

// DeclineFriendRequestHandler godoc
// @Summary Отклонить запрос на добавление в друзья
// @Description Отклонить запрос на добавление в друзья от пользователя с указанным логином. Требуется JWT токен в заголовке.
// @Tags Friends
// @Accept json
// @Produce json
// @Param friend_login query string true "Логин друга"
// @Success 200 {object} map[string]string {"message": "Запрос на добавление в друзья отклонен"}
// @Failure 400 {string} string "Логин друга обязателен"
// @Failure 404 {string} string "Друг не найден"
// @Failure 500 {string} string "Ошибка при отклонении запроса"
// @Security BearerAuth
// @Router /api/friends/decline [post]
func DeclineFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем UUID текущего пользователя из контекста
	userID := r.Context().Value(middleware.UserUUIDKey).(string)

	// Получаем friendID из параметров запроса
	friendLogin := r.URL.Query().Get("friend_login")
	if friendLogin == "" {
		http.Error(w, "Логин друга обязателен", http.StatusBadRequest)
		return
	}

	// Получаем информацию о друге
	friendUser, err := service.GetUserByLogin(friendLogin)
	if err != nil {
		http.Error(w, "Друг не найден", http.StatusNotFound)
		return
	}

	// Отклоняем запрос на добавление в друзья
	err = declineFriendRequest(userID, friendUser.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ответ клиенту
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Запрос на добавление в друзья отклонен"})
}

// RemoveFriendRequestHandler godoc
// @Summary Удалить друга
// @Description Удалить пользователя из списка друзей. Требуется JWT токен в заголовке.
// @Tags Friends
// @Accept json
// @Produce json
// @Param friend_login query string true "Логин друга"
// @Success 200 {object} map[string]string {"message": "Друг успешно удален"}
// @Failure 400 {string} string "Логин друга обязателен"
// @Failure 404 {string} string "Друг не найден"
// @Failure 500 {string} string "Ошибка при удалении друга"
// @Security BearerAuth
// @Router /api/friends/remove [post]
func RemoveFriendRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем UUID текущего пользователя из контекста
	userID := r.Context().Value(middleware.UserUUIDKey).(string)

	// Получаем friendLogin из параметров запроса
	friendLogin := r.URL.Query().Get("friend_login")
	if friendLogin == "" {
		http.Error(w, "Логин друга обязателен", http.StatusBadRequest)
		return
	}

	// Получаем информацию о друге
	friendUser, err := service.GetUserByLogin(friendLogin)
	if err != nil {
		http.Error(w, "Друг не найден", http.StatusNotFound)
		return
	}

	// Удаляем запись о дружбе
	err = removeFriend(userID, friendUser.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ответ клиенту
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Друг успешно удален"})
}
