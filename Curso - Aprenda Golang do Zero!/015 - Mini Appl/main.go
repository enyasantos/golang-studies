package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	appl := app.Gerar()
	if error := appl.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}

//run: go run main.go ip --host amazon.com.br
//executable: ./linha-de-comando ip --host amazon.com.br
