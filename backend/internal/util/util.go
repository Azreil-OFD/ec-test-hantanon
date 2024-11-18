package util

import (
    "golang.org/x/crypto/bcrypt"
    "log"
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
    // Сравниваем пароль с хешем
    err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
    if err != nil {
        log.Println("Ошибка сравнения паролей:", err)
        return false
    }
    return true
}