# Используем официальный образ Go
FROM golang:1.23.3-alpine3.20 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Собираем приложение
RUN go build -o main cmd/main.go

# Создаем финальный образ
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранное приложение из образа builder
COPY --from=builder /app/main .

# Устанавливаем переменные окружения
ENV POSTGRESQL_HOST=195.133.50.12
ENV POSTGRESQL_PORT=5432
ENV POSTGRESQL_USER=gen_user
ENV POSTGRESQL_PASSWORD=12345678A
ENV POSTGRESQL_NAME=default_db
ENV JWT_SECRET_KEY=пушкабомба!

# Команда для запуска приложения
CMD ["./main"]
