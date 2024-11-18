package routs

import (
    "backend/internal/registration" // Путь до вашего пакета регистрации
    "log"
    "net/http"
)

// RegisterRoutes регистрирует все роуты для вашего приложения
func RegisterRoutes() {
    // Регистрируем обработчик для регистрации пользователей
    http.HandleFunc("/api/register", registration.RegisterHandler)

    // Вы можете добавить другие роуты здесь, например:
    // http.HandleFunc("/login", login.LoginHandler)

    log.Println("Server started on :8000...")

    // Запуск сервера
    if err := http.ListenAndServe(":8000", nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}