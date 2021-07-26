package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("executando...")
	log.Fatal(http.ListenAndServe(":3005", nil))
}
