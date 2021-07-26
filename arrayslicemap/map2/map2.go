package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose": 123.123,
		"Joao": 321.321,
		"Pedro": 245.252,
	}

	funcsESalarios["Jonas"] = 949.454
	delete(funcsESalarios, "inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
