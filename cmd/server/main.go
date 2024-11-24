package main

import (
	"log"
	"net/http"

	"github.com/1ssk/Cloud.git/internal/db"
	"github.com/1ssk/Cloud.git/internal/handlers"
)

func main() {
	// Инициализация базы данных
	err := db.Init()
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}
	defer db.Close()

	// Настройка маршрутов
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	http.HandleFunc("/api/files", handlers.UploadFile)
	http.HandleFunc("/api/files/list", handlers.ListFiles)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html") // Главная страница
	})

	// Запуск сервера
	log.Println("Сервер запущен на порту :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}
}
