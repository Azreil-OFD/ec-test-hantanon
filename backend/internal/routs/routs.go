package routs

import (
	"backend/internal/auth"         // Путь до пакета авторизации
	"backend/internal/friends"      // Путь до пакета с друзьями
	"backend/internal/middleware"   // Путь до пакета middleware
	"backend/internal/profile"      // Путь до пакета профиля
	"backend/internal/registration" // Путь до пакета регистрации
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

// RegisterRoutes регистрирует все роуты для вашего приложения
func RegisterRoutes() {
	// Создаем новый мультиплексор
	mux := http.NewServeMux()

	// Добавляем все маршруты с Swagger
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Регистрируем маршруты без TokenAuthMiddleware (только для регистрации и авторизации)
	mux.Handle("/api/register", middleware.NoCORSHandler(http.HandlerFunc(registration.RegisterHandler)))
	mux.Handle("/api/auth", middleware.NoCORSHandler(http.HandlerFunc(auth.LoginHandler)))

	// Регистрируем маршруты с TokenAuthMiddleware и NoCORSHandler для остальных API
	routes := []struct {
		path    string
		handler http.Handler
	}{
		{"/api/profile", http.HandlerFunc(profile.GetProfile)},
		{"/api/friends/request", http.HandlerFunc(friends.SendFriendRequestHandler)},
		{"/api/friends/accept", http.HandlerFunc(friends.AcceptFriendRequestHandler)},
		{"/api/friends/decline", http.HandlerFunc(friends.DeclineFriendRequestHandler)},
		{"/api/friends/remove", http.HandlerFunc(friends.RemoveFriendRequestHandler)},
		{"/api/friends", http.HandlerFunc(friends.GetFriendsInfoHandler)},
		{"/api/search", http.HandlerFunc(profile.SearchUserHandler)},
	}

	// Применяем TokenAuthMiddleware и NoCORSHandler ко всем остальным маршрутам
	for _, route := range routes {
		mux.Handle(route.path, middleware.NoCORSHandler(middleware.TokenAuthMiddleware(route.handler)))
	}

	// Запускаем сервер с мультиплексором
	log.Println("Server started on :8000...")

	// Запуск сервера с мультиплексором
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
