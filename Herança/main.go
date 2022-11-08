package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

func (p Pessoa) String() string {
	//return fmt.Sprintf("Nome completo: %s %s \n Idade: %d", p.Nome, p.Sobrenome, p.Idade)
	return fmt.Sprintf("Nome: %s \nIdade: %d", p.Nome, p.Idade)
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string // privado
}

/*
func (p PessoaFisica) String() string {
	//return fmt.Sprintf("Nome completo: %s %s \n Idade: %d", p.Nome, p.Sobrenome, p.Idade)

	return fmt.Sprintf("Nome completo: %s %s \nIdade: %d", p.Pessoa.Nome, p.Sobrenome, p.Pessoa.Idade)
}
*/

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string // privado
}

func main() {
	p := PessoaFisica{
		Pessoa{
			Nome:   "Enya",
			Idade:  23,
			Status: true,
		},
		"Santos",
		"000.000.000-00",
	}
	fmt.Println(p)
}
