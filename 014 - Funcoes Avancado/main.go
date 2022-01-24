package main

import "fmt"

//------------ Funcoes com retorno nomeado ------------
//Sabe qual variavel tem q retorna, e ja instancia essas variaveis
func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

//---------------------------------------------

//------------ Funcoes Variáticas------------
//Uma funcao que pode receber n parametros
//       (recebe de 1 a n int`s)
func soma(numeros ...int) int {
	fmt.Println(numeros) //eh um slice
	total := 0
	for _, valor := range numeros {
		total += valor
	}
	return total
}

//---------------------------------------------

//------------ Funcoes de ponteiro ------------
func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

//---------------------------------------------

//------------ Funcoes de recursivas ------------
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

//---------------------------------------------

// --------------- Funcao Init ----------------
/*
	Funcao que eh executada antes da funcao main.
	Pode ter uma por arquivo.
*/

var n int

func init() {
	n = 10
	fmt.Println("Executando a funcao init")
}

//---------------------------------------------

func main() {
	fmt.Println(n)

	//------------ Funcoes com retorno nomeado ------------
	s, sub := calculos(10, 20)
	fmt.Println(s, sub)

	//---------------------------------------------

	//------------ Funcoes Variáticas------------
	resultado := soma(1, 2, 3, 4, 6, 8, 9, 9, 0, 10, -22, 33)
	fmt.Println(resultado)
	//---------------------------------------------

	//------------ Funcoes de ponteiro ------------
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
	//---------------------------------------------

	//------------ Funcoes recursivas ------------
	fmt.Println("Funções Recursivas")

	posicao := uint(12)

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
	//---------------------------------------------

	//------------ Funcoes Anônimas ------------
	//auto executavel
	retorno := func(texto string, numero int) string {
		fmt.Println(texto)
		return fmt.Sprintf("Recebido -> %s %d.", texto, numero)
	}("Ola Mundo", 10)

	fmt.Println(retorno)
	//---------------------------------------------
}
