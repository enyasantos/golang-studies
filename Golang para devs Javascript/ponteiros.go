package main

func main() {
	data := "Meu dado"

	printPonteiro(&data)
	println(data + "\n")
	data = "Meu dado"
	printData(data)
	println(data + "\n")
}

func printPonteiro(data *string) {
	println("printData")
	println(data)
	println(*data)
	*data = "Dado alterado"
	println("\n-------")
}

func printData(data string) {
	println("printData")
	println(&data)
	println(data)
	data = "Dado nao alterado"
	println("\n-------")
}
