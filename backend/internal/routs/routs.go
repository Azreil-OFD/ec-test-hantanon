package routs

import (
	"backend/internal/auth" // Путь до пакета авторизации
	"backend/internal/friends"
	"backend/internal/middleware"   // Путь до пакета middleware
	"backend/internal/profile"      // Путь до пакета профиля
	"backend/internal/registration" // Путь до пакета регистрации
	"log"
	"net/http"
)

// RegisterRoutes регистрирует все роуты для вашего приложения
func RegisterRoutes() {
	// Регистрируем обработчик для регистрации пользователей
	http.HandleFunc("/api/register", registration.RegisterHandler)

	// Регистрируем обработчик для авторизации (login)
	http.HandleFunc("/api/auth", auth.LoginHandler)

	// Здесь мы используем middleware для проверки JWT
	http.Handle("/api/profile", middleware.TokenAuthMiddleware(http.HandlerFunc(profile.GetProfile)))

	http.Handle("/api/friends/request", middleware.TokenAuthMiddleware(http.HandlerFunc(friends.SendFriendRequestHandler)))

	http.Handle("/api/friends/accept", middleware.TokenAuthMiddleware(http.HandlerFunc(friends.AcceptFriendRequestHandler)))

	http.Handle("/api/friends/decline", middleware.TokenAuthMiddleware(http.HandlerFunc(friends.DeclineFriendRequestHandler)))

	http.Handle("/api/friends/remove", middleware.TokenAuthMiddleware(http.HandlerFunc(friends.RemoveFriendRequestHandler)))

	// http.Handle("/api/users/search", middleware.TokenAuthMiddleware(http.HandlerFunc(profile.SearchUserHandler)))
	http.HandleFunc("/api/search", profile.SearchUserHandler)

	log.Println("Server started on :8000...")

	// Запуск сервера
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
// /api/users/search
