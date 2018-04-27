package main

import (
	"fmt"
)

// aqui usamos uint assim só poderá ser passo o tipo inteiro sem sinal -
func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	fmt.Println(fatorial(5))

	// resultado2 := fatorial(-4)
	// fmt.Println(resultado2)

}
