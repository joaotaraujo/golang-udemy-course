package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select * from usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var id int
		var nome string
		rows.Scan(&id, &nome)
		fmt.Println(id, nome)
	}

}