package main

import (
	"fmt"
)

// interface que sera usada para pessoa e produto
type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// implementa os metodos da interface. Deve implementar todos os metodos da inteface
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

// recebe qualquer coisa do tipo imprimivel
// como imprimivel é uma interface, entao recebe qualquer coisa que implemente a interface imprimivel
func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Lucas", "Mendes"} // aqui faz a implementacao da interface automagica
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.9}
	fmt.Println(coisa.toString())
	imprimir(coisa)

}
