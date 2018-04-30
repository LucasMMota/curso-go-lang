package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}
type carro struct {
}

// Ferrari -------------------------------
type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// precisa ser por referência, pois vai alterar um valor de ferrari
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

// ------ -------------------------------

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// como ferrari.ligarTurbo() vai acessar um ponteiro é necessário que passe por refência para esportivo
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()
	fmt.Println(carro1, carro2)
}
