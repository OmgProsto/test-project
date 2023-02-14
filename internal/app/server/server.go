package server

import (
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Set("Accept", "application/json")
	w.WriteHeader(404)
	w.Write([]byte("Ghbdtn"))
}

func Main() {
	http.HandleFunc("/", sayhello)           // Устанавливаем роутер
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
