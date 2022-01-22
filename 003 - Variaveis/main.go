package main

import "fmt"

//Go eh fortemente tipada
func main() {
	//Tipo explícito
	var str string = "Variavel string explícita"
	fmt.Println(str)

	//Tipo implícito <inferência de tipo>
	str2 := "Variavel string implícita"
	fmt.Println(str2)

	var (
		str3 string = "Exemplo 1"
		str4 string = str3
	)

	fmt.Println(str4)

	str5, str6 := "Exemplo 2", "Exemplo 3"

	fmt.Println(str5, str6)

	const constante string = "Constante 1"
	fmt.Println(constante)

	//Invertendo valores
	str3, str4 = str4, str3
}
