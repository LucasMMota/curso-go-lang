package main

import (
	"fmt"

	"github.com/lucasmmota/area"
)

func main() {
	fmt.Println(area.Circ(6))
	fmt.Println(area.Rect(5, 2))
	// fmt.Println(area._Triangulo(5, 2))
}

// Pacote no meu workspace
/* package area // por padrao mesmo nome da pasta

import (
	"math"
)

// Pi é público pois começa com Maiúscula
const Pi = 3.1416

// Circ é responsável por calcular a área de uma circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a área do retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é visível pois tem _
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura / 2)
}
*/
