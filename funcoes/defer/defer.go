package main

import (
	"fmt"
)

func obterValorAprovado(numero int) int {
	defer fmt.Println("Fim!") // atrasa a execucao da instrucao para a ultima na funcao antes do return

	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}

	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
