package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteracao %s)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("maria", "falaaaa", 3)
	//fale("joao", "aeeee", 1)

	//go fale("joao", "aeeee", 500)
	//go fale("maria", "falaaaa", 500)
	//time.Sleep(time.Second)


	go fale("joao", "aeeee", 10)
	fale("maria", "falaaaa", 5)
}
