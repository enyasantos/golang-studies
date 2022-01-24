package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//if-init -> inicializacao de uma variavel dentro da condicao
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println(outroNumero)
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println(outroNumero)
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println(outroNumero)
		fmt.Println("Entre 0 e -10")
	}

	//fmt.Println(outroNumero) -> ERROR
}
