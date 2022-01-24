package main

import "fmt"

//estrutura de chave e valor nao mutavel
//map[<tipo_da_chave>]<tipo_dos_valores>
func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Enya",
		"sobrenome": "Santos",
	}

	fmt.Println(usuario["nome"])
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Enya",
			"ultimo":   "Santos",
		},
		"curso": {
			"nome":   "Ciencia da Computacao",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Libra",
	}

	fmt.Println(usuario2)
}
