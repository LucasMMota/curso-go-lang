package main

import (
	"fmt"
	"reflect"
)

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	}
	if n.entre(7.0, 8.99) {
		return "B"
	}
	if n.entre(5.0, 7.99) {
		return "C"
	}
	if n.entre(3.0, 4.99) {
		return "D"
	}
	return "E"
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))

	var nota nota = 10
	fmt.Println(reflect.TypeOf(nota), 1.1 <= nota)
}
