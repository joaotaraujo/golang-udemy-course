package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print("Linha")

	fmt.Println("Quebra")
	fmt.Println("Linha")

	x := 3.141516

	//fmt.Println("O valor de x eh " + x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x eh " + xs)
	fmt.Println("O valor de x eh ", x)

	fmt.Printf("O valor de x eh %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
