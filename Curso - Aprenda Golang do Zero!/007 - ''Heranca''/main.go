// "Herança"
package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa    //A "Herança", todos os campos que estao em pessoa tambem vao fazer parte de estudante
	matricula uint32
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Pedro", "Silva", 23, 1.46}
	fmt.Println(p1)

	e1 := estudante{p1, 123123, "Engenharia", "ITA"}
	fmt.Println(e1.nome)
	fmt.Println(e1)
}
