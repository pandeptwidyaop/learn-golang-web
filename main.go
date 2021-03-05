package main

import (
	"log"
	"net/http"
	"web/handlers"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	log.Println("Application running on port 8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
