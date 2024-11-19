package auth

import (
	"backend/internal/database"
	"backend/internal/util"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Структура для хранения данных пользователя, получаемых из запроса
type user struct {
	uuid     string
	Login    string `json:"login"`    // Логин пользователя
	Password string `json:"password"` // Пароль пользователя
}

// Функция для обработки запроса авторизации
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из тела запроса (login и password)
	var user user
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		// Если не удается распарсить JSON из запроса, возвращаем ошибку
		http.Error(w, "Неверное тело запроса", http.StatusBadRequest) // Русский текст ошибки
		return
	}

	// Проверяем, существует ли пользователь с данным логином в базе данных
	dbUser, err := getUserByLogin(user.Login)
	if err != nil {
		// Если пользователь не найден или произошла ошибка, возвращаем ошибку авторизации
		http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized) // Русский текст ошибки
		return
	}

	// Сравниваем хешированный пароль, который хранится в базе данных, с паролем, введенным пользователем
	if !util.ComparePassword(dbUser.Password, user.Password) {
		// Если пароли не совпадают, возвращаем ошибку авторизации
		http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized) // Русский текст ошибки
		return
	}

	// Если авторизация прошла успешно, генерируем JWT токен для дальнейшей аутентификации
	token, err := util.GenerateJWT(dbUser.uuid)
	if err != nil {
		// Если не удается сгенерировать токен, возвращаем ошибку
		http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError) // Русский текст ошибки
		return
	}

	// Отправляем токен обратно пользователю в формате JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token, // Возвращаем JWT токен
	})
}

// Функция для получения пользователя из базы данных по логину
func getUserByLogin(login string) (*user, error) {
	// SQL запрос для получения данных пользователя по логину
	query := `SELECT id, login, password FROM users WHERE login = $1`
	row := database.DB.QueryRow(context.Background(), query, login)

	// Создаем объект User для хранения полученных данных
	var user user
	err := row.Scan(&user.uuid, &user.Login, &user.Password)
	if err != nil {
		// Если пользователь не найден в базе данных
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("пользователь не найден") // Возвращаем ошибку, что пользователь не найден
		}
		// Если произошла ошибка при выполнении запроса
		log.Println("Ошибка при получении пользователя из базы данных:", err)
		return nil, err
	}
	// Если пользователь найден, возвращаем данные
	return &user, nil
}
