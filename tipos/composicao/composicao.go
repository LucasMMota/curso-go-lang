package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo // pega todos o métodos de esportivo
	luxuoso   // pega todos o métodos de luxuoso
	// pode adicionar mais ...
}

type bwm7 struct{}

func (b bwm7) ligarTurbo() {
	fmt.Println("Turbo ...")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
