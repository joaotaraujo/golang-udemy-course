package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)  //cria numero aleatorio de acordo com milliseconds de agora
}

func main() {
	if i := numeroAleatorio(); i > 5 {  //tbm eh suportado no switch
		fmt.Println("Perdeu")
	} else {
		fmt.Println("Ganhou")
	}
}
