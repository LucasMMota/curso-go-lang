package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[12344444] = "Pedro"
	aprovados[12323233] = "Ana"
	aprovados[12332333] = ""

	fmt.Println(aprovados)

	for cpf, nome := range aprovados { // for key value
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println("")
	fmt.Println(aprovados[12344444])
	delete(aprovados, 12344444) // remove do map
	fmt.Println(aprovados[12344444])

}
