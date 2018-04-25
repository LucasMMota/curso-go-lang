package main

import (
	"fmt"
	m "math" // é possível renomear o packote importado com o novo nome antes da string
)

func main() {
	// go é fortemente tipado, uma vez a var de um tipo, ela não muda

	const PI float64 = 3.1415

	// o compilador pode inferir o tipo
	var raio = 3.2 // tipo float64

	// forma reduzida de criar uma var, utilizando :=
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é:", area)

	const (
		a = 1
		b = 2
	)

	// se nao usar vars declaradas é gerado um erro de compilação
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
