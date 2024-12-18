package middleware

import (
	"backend/internal/util"
	"context"
	"net/http"
	"strings"
)

// Ключ для хранения данных в контексте
type contextKey string

const UserUUIDKey contextKey = "userUUID"

func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем JWT токен из заголовков запроса
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "Токен не предоставлен", http.StatusUnauthorized)
			return
		}

		// Убираем префикс "Bearer " из токена
		if strings.HasPrefix(tokenStr, "Bearer ") {
			tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		} else {
			http.Error(w, "Неверный формат токена", http.StatusUnauthorized)
			return
		}

		// Извлекаем UUID из токена
		uuid, err := util.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}

		// Сохраняем UUID в контексте запроса
		ctx := context.WithValue(r.Context(), UserUUIDKey, uuid)

		// Передаем управление следующему обработчику с обновленным контекстом
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NoCORSHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Разрешаем все домены
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        
        // Обработка preflight запросов (OPTIONS)
        if r.Method == "OPTIONS" {
            return
        }
        
        next.ServeHTTP(w, r)
    })
}