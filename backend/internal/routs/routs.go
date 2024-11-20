package routs

import (
	"backend/internal/auth" // Путь до пакета авторизации
	"backend/internal/friends"
	"backend/internal/middleware"   // Путь до пакета middleware
	"backend/internal/profile"      // Путь до пакета профиля
	"backend/internal/registration" // Путь до пакета регистрации
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// RegisterRoutes регистрирует все роуты для вашего приложения
func RegisterRoutes() {
	// Создаем новый мультиплексор
	mux := http.NewServeMux()

	// Добавляем все маршруты в мультиплексор
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Регистрируем обработчики для всех API с отключенным CORS
	mux.Handle("/api/register", middleware.NoCORSHandler(http.HandlerFunc(registration.RegisterHandler)))
	mux.Handle("/api/auth", middleware.NoCORSHandler(http.HandlerFunc(auth.LoginHandler)))

	// Используем middleware для проверки JWT и отключения CORS
	mux.Handle("/api/profile", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(profile.GetProfile))))
	mux.Handle("/api/friends/request", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(friends.SendFriendRequestHandler))))
	mux.Handle("/api/friends/accept", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(friends.AcceptFriendRequestHandler))))
	mux.Handle("/api/friends/decline", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(friends.DeclineFriendRequestHandler))))
	mux.Handle("/api/friends/remove", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(friends.RemoveFriendRequestHandler))))
	mux.Handle("/api/friends", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(profile.GetFriendsInfoHandler))))
	mux.Handle("/api/search", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(profile.SearchUserHandler))))

	// Запускаем сервер с мультиплексором
	log.Println("Server started on :8000...")

	// Запуск сервера с мультиплексором
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
