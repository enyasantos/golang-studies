package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func alunoEstaAprovado(notas ...float32) bool {
	defer fmt.Println("Media calculada. Resultado sera retornado!") //executa imediatamente antes do return
	var soma float32 = 0.0
	for _, nota := range notas {
		soma += nota
	}

	return (soma / float32(len(notas))) >= 6.0
}

func main() {
	defer funcao1()
	//DEFER = ADIAR (adiar ate o ultimo momento possivel)
	funcao2()

	estaAprovado := alunoEstaAprovado(6.0, 6.0)
	fmt.Println(estaAprovado)
}
