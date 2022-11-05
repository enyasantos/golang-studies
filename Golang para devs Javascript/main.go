package main

import (
	"errors"
	"fmt"
)

func main() {
	var name string = "Enya"
	var age int16
	lastName := "Santos"
	age = 22
	fullName, err, _ := printMe(name, lastName, age)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(fullName)

	result, err := dividir(10, -2)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(result)

	//Map
	mapGo := make(map[string]string)
	mapGo["cidade"] = "Mariana"
	mapGo["estado"] = "Minas Gerais"

	fmt.Println(mapGo)

	//Slice & Array
	//Array
	var valores1 [2]int

	fmt.Println(valores1)

	//Slice (Lista encadeada)
	var valores2 []int
	fmt.Println(valores2)
}

func dividir(value1 int, value2 int) (int, error) {
	if value2 == 0 {
		return 0, errors.New("O segundo valor deve ser diferente de zero")
	}

	return value1 / value2, nil
}

func printMe(name string, lastname string, age int16) (string, error, int16) {
	return name + " " + lastname + " tem " + string(rune(age)) + " anos", nil, age
}
