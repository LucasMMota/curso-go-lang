package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:@/curso_go")
	if err != nil {
		panic(err)
	}
	defer db.Close() // atrasa o fechamento da conexão

	rows, _ := db.Query("select * from usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome) // mapea para o result set do db para o struct
		fmt.Println(u)

	}
}
