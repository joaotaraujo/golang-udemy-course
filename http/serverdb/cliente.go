package main

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Usuario struct {
	ID int `json:"id"`
	Nome string `json:"nome"`
}

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "desculpa... :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var c Usuario
	db.QueryRow("select id, nome from usuarios where id > ?", id).Scan(&c.ID, &c.Nome)

	json, _ := json2.Marshal(c)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u Usuario
	rows, _ := db.Query("select * from usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&u.ID, &u.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json2.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}