package service

import (
	"backend/internal/database"
	"context"
	"fmt"
	"log"
)

type User struct {
	UUID     string
	FullName string
	Login    string
	Email    string
}

func GetUserByLogin(login string) (*User, error) {
	// SQL запрос для получения данных пользователя по UUID
	query := `SELECT id, login, email, full_name FROM users WHERE login = $1`
	row := database.DB.QueryRow(context.Background(), query, login)

	// Создаем объект UserProfile для хранения полученных данных
	var user User
	err := row.Scan(&user.UUID, &user.Login, &user.Email, &user.FullName)
	if err != nil {
		// Если пользователь не найден в базе данных
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("пользователь не найден")
		}
		log.Println("Ошибка при получении пользователя из базы данных:", err)
		return nil, err
	}

	// Возвращаем профиль пользователя
	return &user, nil
}

// Функция для получения пользователя из базы данных по UUID
func GetUserByUUID(uuid string) (*User, error) {
	// SQL запрос для получения данных пользователя по UUID
	query := `SELECT id, login, email, full_name FROM users WHERE id = $1`
	row := database.DB.QueryRow(context.Background(), query, uuid)

	// Создаем объект UserProfile для хранения полученных данных
	var user User
	err := row.Scan(&user.UUID, &user.Login, &user.Email, &user.FullName)
	if err != nil {
		// Если пользователь не найден в базе данных
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("пользователь не найден")
		}
		log.Println("Ошибка при получении пользователя из базы данных:", err)
		return nil, err
	}

	// Возвращаем профиль пользователя
	return &user, nil
}
