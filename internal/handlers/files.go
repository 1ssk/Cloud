package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFile загружает файл на сервер
func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка получения файла: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Создание директории, если её нет
	os.MkdirAll("./uploads", os.ModePerm)

	// Создание нового файла
	dst, err := os.Create(filepath.Join("./uploads", header.Filename))
	if err != nil {
		http.Error(w, "Ошибка сохранения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Копирование содержимого файла
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Ошибка сохранения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Файл успешно загружен"))
}

// ListFiles возвращает список загруженных файлов
func ListFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	files, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, "Ошибка чтения директории: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileNames)
}
