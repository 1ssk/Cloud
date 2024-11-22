package handlers

import (
	"net/http"
	"text/template"
)

func CloudHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/cloud.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
