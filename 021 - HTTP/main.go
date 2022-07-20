package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar dados de usuários!"))
}

func main() {
	http.HandleFunc("/", home)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
