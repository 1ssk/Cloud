package main

import (
	"fmt"
	"net/http"

	"github.com/1ssk/Cloud.git/app/handlers"
)

func main() {
	// Обработка папки /static/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.CloudHandler)
	fmt.Println("Server starting on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
