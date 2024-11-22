package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/1ssk/Cloud.git/handlers"
)

func main() {
	// Обработчики маршрутов
	http.HandleFunc("/", handlers.CloudHandler)          // Главная страница с формой загрузки
	http.HandleFunc("/upload", handlers.UploadHandler)   // Обработка загрузки файлов
	http.HandleFunc("/files", handlers.ListFilesHandler) // Обработка списка файлов

	// Обработчик для скачивания файлов из папки uploads
	http.HandleFunc("/uploads/", func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем имя файла из URL (удаляем "/uploads/" из начала пути)
		fileName := r.URL.Path[len("/uploads/"):]
		filePath := fmt.Sprintf("%s/%s", handlers.UploadDir, fileName) // Используем переменную UploadDir из handlers

		// Открываем файл для чтения
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "Не удалось найти файл", http.StatusNotFound)
			return
		}
		defer file.Close()

		// Устанавливаем заголовок контента для скачивания
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

		// Копируем содержимое файла в ответ
		http.ServeFile(w, r, filePath)
	})

	// Запуск сервера
	log.Println("Сервер запущен на порту :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
