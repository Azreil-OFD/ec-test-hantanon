package friends

import (
	"backend/internal/database"
	"context"
	"errors"
	"log"
)

const (
	Pending  = "pending"
	Accepted = "accepted"
	Declined = "declined"
)

// Отправка запроса на добавление в друзья
func sendFriendRequest(userID, friendID string) error {
	// Проверяем, что пользователь не пытается добавить себя в друзья
	if userID == friendID {
		return errors.New("нельзя добавить себя в друзья")
	}

	// Проверяем, существует ли уже запрос на добавление в друзья от текущего пользователя
	query := `SELECT status FROM friends WHERE user_id = $1 AND friend_id = $2`
	row := database.DB.QueryRow(context.Background(), query, userID, friendID)

	var status string
	err := row.Scan(&status)

	// Если запрос уже существует
	if err == nil {
		// Если статус запроса "pending" — запрос уже был отправлен
		if status == Pending {
			return errors.New("запрос уже отправлен")
		} else if status == Accepted {
			// Если статус "accepted" — вы уже друзья
			return errors.New("вы уже друзья")
		} else if status == Declined {
			// Если запрос был отклонен — нельзя снова отправить запрос
			return errors.New("вы не можете отправить запрос, так как он был отклонен")
		}
	}

	// Проверяем, отправил ли друг вам запрос
	reverseQuery := `SELECT status FROM friends WHERE user_id = $1 AND friend_id = $2`
	row = database.DB.QueryRow(context.Background(), reverseQuery, friendID, userID)

	err = row.Scan(&status)
	// Если запрос был отправлен другом, но ещё не принят
	if err == nil {
		if status == Pending {
			return errors.New("ваш друг отправил вам запрос, примите его")
		} else if status == Accepted {
			return errors.New("вы уже друзья")
		}
	}

	// Если нет существующего запроса с вашим другом, отправляем новый запрос
	insertQuery := `INSERT INTO friends (user_id, friend_id, status) VALUES ($1, $2, $3)`
	_, err = database.DB.Exec(context.Background(), insertQuery, userID, friendID, Pending)
	if err != nil {
		log.Println("Ошибка при отправке запроса:", err)
		return err
	}
	return nil
}

func acceptFriendRequest(userID, friendID string) error {
	// Проверяем, существует ли запрос в ожидании
	query := `SELECT status FROM friends WHERE user_id = $1 AND friend_id = $2`
	row := database.DB.QueryRow(context.Background(), query, friendID, userID)

	var status string
	err := row.Scan(&status)
	if err != nil {
		return errors.New("запрос на добавление в друзья не найден")
	}

	if status != Pending {
		return errors.New("нельзя принять запрос в этом состоянии")
	}

	// Обновляем статус на 'accepted'
	updateQuery := `UPDATE friends SET status = $1 WHERE user_id = $2 AND friend_id = $3`
	_, err = database.DB.Exec(context.Background(), updateQuery, Accepted, friendID, userID)
	if err != nil {
		log.Println("Ошибка при принятии запроса:", err)
		return err
	}

	insertQuery := `
    INSERT INTO friends (user_id, friend_id, status)
    VALUES ($1, $2, $3)
    ON CONFLICT (user_id, friend_id) DO UPDATE
    SET status = EXCLUDED.status
	`
	_, err = database.DB.Exec(context.Background(), insertQuery, userID, friendID, Accepted)
	if err != nil {
		log.Println("Ошибка при принятии обратного запроса:", err)
		return err
	}

	return nil
}

// Отклонение запроса на добавление в друзья
func declineFriendRequest(userID, friendID string) error {
	// Проверяем, существует ли запрос в ожидании
	query := `SELECT status FROM friends WHERE user_id = $1 AND friend_id = $2`
	row := database.DB.QueryRow(context.Background(), query, friendID, userID)

	var status string
	err := row.Scan(&status)
	if err != nil {
		return errors.New("запрос на добавление в друзья не найден")
	}

	// Проверяем, что статус запроса в ожидании
	if status != Pending {
		return errors.New("нельзя отклонить запрос в этом состоянии")
	}

	// Обновляем статус на "declined" вместо удаления записи
	updateQuery := `UPDATE friends SET status = $1 WHERE user_id = $2 AND friend_id = $3`
	_, err = database.DB.Exec(context.Background(), updateQuery, Declined, friendID, userID)
	if err != nil {
		log.Println("Ошибка при обновлении статуса запроса:", err)
		return err
	}

	// // Можно обновить статус и для обратной записи, если хотите, чтобы обе стороны видели отклонение
	// _, err = database.DB.Exec(context.Background(), updateQuery, Declined, friendID, userID)
	// if err != nil {
	// 	log.Println("Ошибка при обновлении статуса обратного запроса:", err)
	// 	return err
	// }

	return nil
}

// Удаление пользователя из списка друзей
func removeFriend(userID, friendID string) error {
	// Удаляем запись о дружбе из таблицы
	deleteQuery := `DELETE FROM friends WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)`
	_, err := database.DB.Exec(context.Background(), deleteQuery, userID, friendID)
	if err != nil {
		log.Println("Ошибка при удалении из друзей:", err)
		return err
	}

	return nil
}
