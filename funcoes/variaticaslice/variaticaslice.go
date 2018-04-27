package main

import (
	"fmt"
)

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"} // crio um slice
	imprimirAprovados(aprovados...)                                // posso passar um slice com ... e ele entende que é a lista de argumentos
}
