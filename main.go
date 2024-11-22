package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/files/", serveFile)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the file storage service</h1>"+
		"<form action='/upload' method='post' enctype='multipart/form-data'>"+
		"<input type='file' name='file' id='file'>"+
		"<input type='submit' value='Upload'>"+
		"</form>")
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Parse up to 10 MB of files

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving file from form data")
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("uploads", handler.Filename)
	if err != nil {
		fmt.Println("Error creating temporary file")
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "File uploaded successfully\n")
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path[len("/files/"):]
	file, err := os.Open("uploads/" + fileName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Error getting file info", http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, fileName, fileInfo.ModTime(), file)
}
