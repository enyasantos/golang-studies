package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Ola mundo", canal)

	//mensagem, open := <-canal //canal vai esperar receber o valor

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto // mandando um valor para o canal
		time.Sleep(time.Second)
	}

	close(canal)
}
