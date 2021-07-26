package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
	Tags []string `json:"tags"`
}

func main() {
	// struct to json
	produto1 := produto{1, "notebook", 1999.90, []string{"Promocao", "Eletronico"}}
	produto1Json, _ := json.Marshal(produto1)
	fmt.Println(string(produto1Json))

	// json to struct
	var p2 produto
	jsonString := `{"id":2,"nome":"aeee","preco":8.9,"tags":["papelaria","importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])

}
