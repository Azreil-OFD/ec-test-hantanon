package routs

import (
	_ "backend/cmd/docs"
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
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	// Регистрируем обработчик для регистрации пользователей
	http.Handle("/api/register", middleware.NoCORSHandler(http.HandlerFunc(registration.RegisterHandler)))

	// Регистрируем обработчик для авторизации (login)
	http.Handle("/api/auth", middleware.NoCORSHandler(http.HandlerFunc(auth.LoginHandler)))

	// Здесь мы используем middleware для проверки JWT
	http.Handle("/api/profile", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(profile.GetProfile))))

	http.Handle("/api/friends/request", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(friends.SendFriendRequestHandler))))

	http.Handle("/api/friends/accept", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(friends.AcceptFriendRequestHandler))))

	http.Handle("/api/friends/decline", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(friends.DeclineFriendRequestHandler))))

	http.Handle("/api/friends/remove", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(friends.RemoveFriendRequestHandler))))

	http.Handle("/api/friends", middleware.NoCORSHandler(middleware.TokenAuthMiddleware((http.HandlerFunc(profile.GetFriendsInfoHandler)))))

	http.Handle("/api/search", middleware.NoCORSHandler(middleware.TokenAuthMiddleware(http.HandlerFunc(profile.SearchUserHandler))))

	log.Println("Server started on :8000...")

	// Запуск сервера
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// /api/users/search
// http.HandleFunc("/api/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
// 	// Отдаём файл openapi.yaml или openapi.json
// 	http.ServeFile(w, r, "./openapi.yaml") // Убедитесь, что файл существует по этому пути
// })
