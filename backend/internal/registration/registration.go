package registration

import (
	"backend/internal/database"
	"backend/internal/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type user struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Только POST-запросы
	if r.Method != http.MethodPost {
		http.Error(w, "Неверный метод запроса", http.StatusMethodNotAllowed)
		return
	}

	var user user

	// Декодируем JSON тело запроса
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest)
		return
	}

	// Простая валидация входных данных
	if strings.TrimSpace(user.Login) == "" || strings.TrimSpace(user.Password) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.FullName) == "" {
		http.Error(w, "Логин, пароль, email и имя обязательны", http.StatusBadRequest)
		return
	}

	// Здесь можно добавить дополнительные проверки (например, проверка на уникальность логина или email)

	// Хешируем пароль перед сохранением в базу данных (рекомендуется использовать более безопасные хеш-функции, например bcrypt)
	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Ошибка при хешировании пароля", http.StatusInternalServerError)
		return
	}

	// Добавляем пользователя в базу данных
	err = saveUserToDB(user)
	if err != nil {
		http.Error(w, "Такой пользователь уже существует", http.StatusConflict)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Пользователь %s успешно зарегистрирован!", user.Login)
}

// Функция для сохранения пользователя в базе данных
func saveUserToDB(user user) error {
	// Пример запроса для вставки пользователя
	query := `INSERT INTO users (login, password, email, full_name) VALUES ($1, $2, $3, $4)`
	_, err := database.DB.Exec(context.Background(), query, user.Login, user.Password, user.Email, user.FullName)
	if err != nil {
		// Логируем ошибку на русском
		log.Println("Ошибка при вставке пользователя в базу данных:", err)
		return err
	}

	return nil
}
