package main

import (
	"fmt"
	"time"
)

func main() {
	// ConcorrĂȘncia != Paralelismo
	go texto("Mensagem 1")
	texto("Mensagem 2")

}

func texto(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
