package models

// File представляет файл, загруженный на сервер
type File struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}
