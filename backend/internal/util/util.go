package util

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword хеширует пароль с использованием bcrypt
func HashPassword(password string) (string, error) {
	// Генерируем хеш с уровнем сложности 12 (можно увеличить для повышения безопасности)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword проверяет, совпадает ли введённый пароль с сохранённым хешем
func ComparePassword(storedHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

var secretKey = os.Getenv("JWT_SECRET_KEY")

func GenerateJWT(userID string) (string, error) {
	// Определяем время истечения токена (например, 12 часов)
	tokenExpirationTime := time.Now().Add(12 * time.Hour)

	// Создаем claims с добавлением читаемого времени и типа токена
	claims := jwt.MapClaims{
		"uuid": userID,                     // ID пользователя
		"exp":  tokenExpirationTime.Unix(), // Время истечения в Unix формате (обязательно для JWT)
	}
	// Создаем новый токен с алгоритмом подписи и claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с помощью секретного ключа
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Валидация JWT с использованием конфигурации
func ValidateJWT(tokenString string) (string, error) {
	// Парсим и валидируем токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что метод подписи токена это HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неожиданный метод подписи")
		}
		return []byte(secretKey), nil
	})

	// Если произошла ошибка при парсинге токена
	if err != nil {
		return "", err
	}

	// Если токен валиден, извлекаем данные из claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Извлекаем UUID из claims и возвращаем его
		if uuid, ok := claims["uuid"].(string); ok {
			return uuid, nil
		}
	}

	// Если UUID не найден в claims или токен невалиден
	return "", errors.New("недействительный токен или отсутствует UUID")
}
