package main

import (
	"fmt"
)

// ... -> pode receber vários parâmetros
func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros { //numeros vem como um array do tipo passado
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
