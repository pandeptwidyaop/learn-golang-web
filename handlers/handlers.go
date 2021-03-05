package handlers

import (
	"log"
	"net/http"
	"path"
	"text/template"
)

//Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	temp, err := template.ParseFiles(path.Join("views", "home.html"))
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Ohh no", http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"title":   "Pande Ganteng",
		"address": "Jalan Sesame",
	}

	err = temp.Execute(w, data)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Ohh no", http.StatusInternalServerError)
		return
	}
}
