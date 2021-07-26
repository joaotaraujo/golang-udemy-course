package main

import "fmt"

func main() {
	i := 1

	// go nao tem aritmetica de ponteiros
	var p *int = nil

	p = &i
	*p++
	i++

	fmt.Println(p, *p, i, &i)
}
