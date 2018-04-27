package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

// precisa passar como ponteiro para alterar a variavel
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ") // qubro a string em um array
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Lucas", "Mendes"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Lorde Lucas")
	fmt.Println(p1.getNomeCompleto())
}
