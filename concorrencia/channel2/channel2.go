package main

import (
	"fmt"
	"time"
)

// channel (canal) - eh a forma de comunicacao entra goroutines
// channel eh um tipo de dado

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base  // envia dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base  // envia dados para o canal

	time.Sleep(3 * time.Second)
	c <- 4 * base  // envia dados para o canal
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c  //recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
