package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("enyalgs@gmail.com")
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println("Email valido")
	}
}
