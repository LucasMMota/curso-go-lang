package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4 // float64
	y := 2   // int

	fmt.Println(x / float64(y)) // é necessário sempre converter os tipos para operar

	nota := 6.9
	notaFinal := int(nota) // pega a parte inteira
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97)) // string() converte o int para o valor correspondente na tabela ascii

	// int para string
	fmt.Println("teste " + strconv.Itoa(123))

	// converte a string 123 para int
	num, _ := strconv.Atoi("123")
	fmt.Println("O valor é " + strconv.Itoa(num))
	fmt.Println(num - 122)

	// converte a stringe true para bool
	b, _ := strconv.ParseBool("true") // bool true pode ser "true" ou "1"
	if b {
		fmt.Println("Verdadeiro")
	}

}
