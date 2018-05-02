package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // esta com _ pq o lint nao acha a referencia automaticamente na linha 18 e remove
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:@/") // cria a conexao, usa indiretamente github.com/go-sql-driver/mysql
	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists curso_go")
	exec(db, "use curso_go")
	exec(db, "drop table if exists usuario")
	exec(db, ` create table usuarios (
			id integer auto_increment,
			nome varchar(80),
			PRIMARY KEY (id)
		)`)

}
