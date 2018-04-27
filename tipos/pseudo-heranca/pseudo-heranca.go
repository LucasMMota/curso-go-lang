package main

import (
	"fmt"
)

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos, tudo que é de carro estará disponível dentro de ferrari. Faz uma composição
	turboLigado bool
}

func main() {
	f := ferrari{} // define a estrutura do tipo ferrari vazia
	f.nome = "F40" // acessa os campos do tipo carro dentro de ferrari
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f)
}
