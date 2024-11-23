package friends

import (
	"backend/internal/database"
	"backend/internal/middleware"
	"backend/internal/model"
	"context"
	"fmt"
	"log"
	"net/http"
)

func GetFriendsInfoHandler(w http.ResponseWriter, r *http.Request) {
	response := model.Response{}

	userID := r.Context().Value(middleware.UserUUIDKey).(string)

	requestType := r.URL.Query()["type"]

	result, err := getFriendInfo(userID, requestType)
	if err != nil {
		customErr, _ := err.(*model.CustomError)
		response.Message = customErr.Message
		response.Status = customErr.Code
		model.SendJSONResponse(w, response)
		return
	}

	response.Message = "Запрос на получение друзей прошла успешно!"
	response.Status = http.StatusOK
	response.Data = result
	model.SendJSONResponse(w, response)
}

func getFriendInfo(userID string, requestTypes []string) (map[string][]model.User, error) {
	result := make(map[string][]model.User)
	if len(requestTypes) == 0 {
		friends, err := getFriends(userID)
		if err != nil {
			return nil, err
		}
		incoming, err := getIncomingRequests(userID)
		if err != nil {
			return nil, err
		}
		outgoing, err := getOutgoingRequests(userID)
		if err != nil {
			return nil, err
		}
		result["friends"] = friends
		result["incoming"] = incoming
		result["outgoing"] = outgoing
	} else {
		for _, _type_ := range requestTypes {
			switch _type_ {
			case "friends":
				friends, err := getFriends(userID)
				if err != nil {
					return nil, err
				}
				result["friends"] = friends
			case "incoming":
				incoming, err := getIncomingRequests(userID)
				if err != nil {
					return nil, err
				}
				result["incoming"] = incoming
			case "outgoing":
				outgoing, err := getOutgoingRequests(userID)
				if err != nil {
					return nil, err
				}
				result["outgoing"] = outgoing
			default:
				return nil, &model.CustomError{
					Code:    model.BadRequest,
					Message: fmt.Sprintf("Некорректный параметр 'type': %s. Ожидается 'friends', 'incoming' или 'outgoing'", _type_),
				}
			}
		}
	}
	return result, nil
}

// Функция для получения списка друзей
func getFriends(userID string) ([]model.User, error) {
	query := `SELECT u.id, u.login, u.full_name, u.email
			  FROM users u
			  JOIN friends f ON u.id = f.friend_id
			  WHERE f.user_id = $1 AND f.status = 'accepted'`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Println("Ошибка: ", err)
		err = &model.CustomError{
			Code:    model.DBError,
			Message: "Ошибка на стороне сервера",
		}
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
			log.Println("Ошибка: ", err)
			err = &model.CustomError{
				Code:    model.DBError,
				Message: "Ошибка на стороне сервера",
			}
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Ошибка: ", err)
		err = &model.CustomError{
			Code:    model.DBError,
			Message: "Ошибка на стороне сервера",
		}
		return nil, err
	}

	return users, nil
}

// Функция для получения списка входящих заявок
func getIncomingRequests(userID string) ([]model.User, error) {
	query := `SELECT u.id, u.login, u.full_name, u.email
			  FROM users u
			  JOIN friends f ON u.id = f.user_id
			  WHERE f.friend_id = $1 AND f.status = 'pending'`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Println("Ошибка: ", err)
		err = &model.CustomError{
			Code:    model.DBError,
			Message: "Ошибка на стороне сервера",
		}
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
			log.Println("Ошибка: ", err)
			err = &model.CustomError{
				Code:    model.DBError,
				Message: "Ошибка на стороне сервера",
			}
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Ошибка: ", err)
		err = &model.CustomError{
			Code:    model.DBError,
			Message: "Ошибка на стороне сервера",
		}
		return nil, err
	}

	return users, nil
}

// Функция для получения списка исходящих заявок
func getOutgoingRequests(userID string) ([]model.User, error) {
	query := `SELECT u.id, u.login, u.full_name, u.email
			  FROM users u
			  JOIN friends f ON u.id = f.friend_id
			  WHERE f.user_id = $1 AND f.status = 'pending'`
	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Println("Ошибка: ", err)
		err = &model.CustomError{
			Code:    model.DBError,
			Message: "Ошибка на стороне сервера",
		}
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UUID, &user.Login, &user.FullName, &user.Email); err != nil {
			log.Println("Ошибка: ", err)
			err = &model.CustomError{
				Code:    model.DBError,
				Message: "Ошибка на стороне сервера",
			}
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Println("Ошибка: ", err)
		err = &model.CustomError{
			Code:    model.DBError,
			Message: "Ошибка на стороне сервера",
		}
		return nil, err
	}

	return users, nil
}
