package service

import (
	"backend/internal/database"
	"backend/internal/model"
	"context"
	"fmt"
	"log"
)

type User struct {
	UUID         string
	FullName     string
	Login        string
	Email        string
	PasswordHash string
}

// *pgx.proxyError
func GetUserByLogin(login string) (*User, error) {
	query := `SELECT id, login, password, email, full_name FROM users WHERE login = $1`
	var user User
	err := database.DB.QueryRow(context.Background(), query, login).Scan(&user.UUID, &user.Login, &user.PasswordHash, &user.Email, &user.FullName)
	if err != nil {
		if err.Error() == "no rows in result set" {
			err = &model.CustomError{
				Message: "Пользователь не найден",
				Code:    model.NotFound,
			}
			return nil, err
		}
		log.Printf("Ошибка: %v", err)
		err = &model.CustomError{
			Message: "Ошибка на стороне сервера ",
			Code:    model.DBError,
		}
		return nil, err
	}
	return &user, nil
}

// func getUserByLogin(login string) (*user, error) {
// 	// SQL запрос для получения данных пользователя по логину
// 	query := `SELECT id, login, password FROM users WHERE login = $1`
// 	var user user
// 	err := database.DB.QueryRow(context.Background(), query, login).Scan(&user.uuid, &user.Login, &user.Password)
// 	if err != nil {
// 		// Если пользователь не найден в базе данных
// 		log.Fatalf("%T", err)
// 		if err.Error() == "sql: no rows in result set" {
// 			return nil, errors.New("пользователь не найден") // Возвращаем ошибку, что пользователь не найден
// 		}
// 		// Если произошла ошибка при выполнении запроса
// 		log.Println("Ошибка при получении пользователя из базы данных:", err)
// 		return nil, err
// 	}
// 	// Если пользователь найден, возвращаем данные
// 	return &user, nil
// }

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
