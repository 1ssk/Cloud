package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
	"time"
)

var UploadDir = "./uploads" // Путь к папке загрузки

// CloudHandler обрабатывает запросы на главную страницу (форму загрузки файла)
func CloudHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/cloud.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

// UploadHandler обрабатывает запросы на загрузку файла
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка, существует ли папка uploads, если нет - создаем
	if _, err := os.Stat(UploadDir); os.IsNotExist(err) {
		err := os.Mkdir(UploadDir, os.ModePerm)
		if err != nil {
			http.Error(w, "Не удалось создать папку для загрузки", http.StatusInternalServerError)
			return
		}
	}

	// Разбор формы для загрузки файлов
	err := r.ParseMultipartForm(10 << 20) // Ограничение по размеру файла 10 MB
	if err != nil {
		http.Error(w, "Ошибка при разборе файла", http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Создаем файл с уникальным именем
	outFile, err := os.Create(fmt.Sprintf("%s/%d_uploaded_file", UploadDir, time.Now().Unix()))
	if err != nil {
		http.Error(w, "Ошибка при сохранении файла", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Копируем содержимое загруженного файла в новый файл
	_, err = outFile.ReadFrom(file)
	if err != nil {
		http.Error(w, "Ошибка при копировании содержимого файла", http.StatusInternalServerError)
		return
	}

	// Отправляем сообщение об успешной загрузке
	fmt.Fprintf(w, "Файл успешно загружен!")
}

// ListFilesHandler отображает список файлов, загруженных в папку uploads
func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	// Читаем список файлов из директории uploads
	files, err := ioutil.ReadDir(UploadDir)
	if err != nil {
		http.Error(w, "Ошибка при чтении директории", http.StatusInternalServerError)
		return
	}

	// Отправляем список файлов в HTML-шаблон
	tmpl, err := template.ParseFiles("templates/list_files.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, files)
}
