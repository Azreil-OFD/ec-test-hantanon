package friends

import (
	"backend/internal/middleware"
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

// Обработчик для отправки запроса на добавление в друзья
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
