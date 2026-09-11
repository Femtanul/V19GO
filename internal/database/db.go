package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB - обёртка над подключением к PostgreSQL
type DB struct {
	conn *sql.DB
}

// Config - настройки подключения
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// NewDB создает новое подключение к PostgreSQL
func NewDB(cfg Config) (*DB, error) {
	// Формируем строку подключения (DSN)
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	// Открываем подключение
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	// Проверяем, что подключение работает
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return &DB{conn: conn}, nil
}

// Close закрывает подключение
func (db *DB) Close() error {
	return db.conn.Close()
}