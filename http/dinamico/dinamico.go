package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2012 03:04:05")
	fmt.Fprintf(w, "<h1>Hora certa: %s<h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("executando...")
	log.Fatal(http.ListenAndServe(":3002", nil))
}
