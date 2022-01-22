package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}

func main() {
	muxDispatcher := mux.NewRouter()

	muxDispatcher.HandleFunc("/users", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode([]User{
			{
				ID:        1,
				Name:      "Maria",
				CreatedAt: "20 jan 2022",
			},
			{
				ID:        2,
				Name:      "Joao",
				CreatedAt: "20 jan 2022",
			},
			{
				ID:        3,
				Name:      "Ana",
				CreatedAt: "20 jan 2022",
			},
		})
	}).Methods("GET")

	port := ":8080"
	fmt.Println("Online on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, muxDispatcher))
}
