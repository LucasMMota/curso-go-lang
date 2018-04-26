package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i //pegando o endereço da variável i e atribuindo em p

	*p++ // dereferenciando (pegando o valor)
	i++

	// GO não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}
