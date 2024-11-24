package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// Init инициализирует подключение к базе данных
func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке файла .env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Printf("Попытка подключения: %s", connStr)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("Ошибка при подключении к базе данных: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("Ошибка при проверке подключения: %v", err)
	}

	log.Println("Подключение к базе данных успешно!")
	return nil
}

// Закрытие соединения с базой данных
func Close() {
	DB.Close()
}
