package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/curso_go")
	if err != nil {
		panic(err)
	}
	defer db.Close() // atrasa o fechamento da conexão

	// update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Hermione da Silva", 1)
	stmt.Exec("Calvo Geraldo", 2)

	//delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}
