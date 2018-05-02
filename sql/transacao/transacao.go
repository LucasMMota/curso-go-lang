package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/curso_go")
	if err != nil {
		panic(err)
	}
	defer db.Close() // atrasa o fechamento da conexão

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values (?,?)")

	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	_, err = stmt.Exec(1, "Thiago") // chave duplicada

	if err != nil {
		tx.Rollback() // volta todos os stmt
		log.Fatal(err)
	}

	tx.Commit()
}
