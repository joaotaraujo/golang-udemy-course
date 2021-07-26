package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Guga": 134.23,
			"Joao": 123.11,
			"Juco": 122.21,
		},
		"B": {
			"Jonas": 1334.23,
			"Padrin": 13423.11,
			"Mano": 12342.21,
		},
	}

	delete(funcsPorLetra, "B")

	for letra, funcs := range funcsPorLetra{
		fmt.Println(letra, funcs)
	}
}

