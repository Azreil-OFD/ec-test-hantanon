package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var DB *pgxpool.Pool

func init() {
	// Загружаем конфигурацию базы данных
	host := os.Getenv("POSTGRESQL_HOST")
	port := os.Getenv("POSTGRESQL_PORT")
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	dbname := os.Getenv("POSTGRESQL_NAME")
	// Создаем строку подключения
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		user, password, host, port, dbname)

	// Настраиваем конфигурацию пула соединений
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Ошибка парсинга конфигурации: %v", err)
	}

	// Настраиваем параметры пула
	config.MaxConns = 10                       // Максимум 10 соединений в пуле
	config.MaxConnLifetime = 30 * time.Minute  // Максимальное время жизни соединения
	config.HealthCheckPeriod = 1 * time.Minute // Период проверки активности соединений

	// Создаем пул с контекстом и тайм-аутом
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Проверяем соединение с помощью Ping
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Ошибка проверки подключения к базе данных (ping): %v", err)
	}

	// Сохраняем пул соединений в глобальной структуре
	DB = pool

	log.Println("Успешное подключение к базе данных с использованием пула соединений")
}
