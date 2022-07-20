package main

import (
	"crud/server"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", server.ListarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.MostrarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", server.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 8080...")
	http.ListenAndServe(":8080", router)
}
