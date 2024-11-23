package service

import (
	"backend/internal/database"
	"backend/internal/model"
	"context"
	"log"
)

// *pgx.proxyError
func GetUserByLogin(login string) (*model.User, error) {
	query := `SELECT id, login, password, email, full_name FROM users WHERE login = $1`
	var user model.User
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
			Message: "Ошибка на стороне сервера",
			Code:    model.DBError,
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByUUID(uuid string) (*model.User, error) {
	query := `SELECT id, login, password, email, full_name FROM users WHERE id = $1`
	var user model.User
	err := database.DB.QueryRow(context.Background(), query, uuid).Scan(&user.UUID, &user.Login, &user.PasswordHash, &user.Email, &user.FullName)
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
			Message: "Ошибка на стороне сервера",
			Code:    model.DBError,
		}
		return nil, err
	}
	return &user, nil
}
