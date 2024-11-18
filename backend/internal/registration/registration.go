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

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Только POST-запросы
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User

	// Декодируем JSON тело запроса
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Простая валидация входных данных
	if strings.TrimSpace(user.Login) == "" || strings.TrimSpace(user.Password) == "" || strings.TrimSpace(user.Email) == "" {
		http.Error(w, "Login, password and email are required", http.StatusBadRequest)
		return
	}

	// Здесь можно добавить дополнительные проверки (например, проверка на уникальность логина или email)

	// Хешируем пароль перед сохранением в базу данных (рекомендуется использовать более безопасные хеш-функции, например bcrypt)
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error saving user to database", http.StatusInternalServerError)
		return
	}

	// Добавляем пользователя в базу данных
	err = saveUserToDB(user.Login, hashedPassword, user.Email)
	if err != nil {
		
		http.Error(w, "Error saving user to database", http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User %s successfully registered!", user.Login)
}

// Функция для хеширования пароля (реализована просто для примера, рекомендуется использовать bcrypt)
func hashPassword(password string) string {
	return password // Здесь нужно заменить на реальную функцию хеширования, например bcrypt
}

// Функция для сохранения пользователя в базе данных
func saveUserToDB(login, password, email string) error {
	// Пример запроса для вставки пользователя
	query := `INSERT INTO users (login, password, email) VALUES ($1, $2, $3)`
	_, err := database.DB.Exec(context.Background(), query, login, password, email)
	if err != nil {
		log.Println("Error inserting user into database:", err)
		return err
	}

	return nil
}
