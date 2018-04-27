package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Dilma da Massa": 15456.78,
			"Inês Brasil":    8769.78,
			"Luke Skywalker": 1000.00,
		},
		"P": {
			"Darth Vader": 4433,
		},
		"J": {
			"Stan Lee": 3333,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Printf("[%s]\n", letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
