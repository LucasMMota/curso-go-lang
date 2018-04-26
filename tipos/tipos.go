package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint 16 uint32 uint 64
	var b byte = 255 //byte é um apelido para uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal.. int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	// rune representa um mapeamento da tabela Unicode
	var i2 rune = 'a'
	fmt.Println("o rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo do literal 49.99 é", reflect.TypeOf(49.99)) // por padrao números reais sao float64

	// boolean
	bo := true
	fmt.Println("o tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Lucas"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome 
	é
	Lucas`
	fmt.Println(s2)
	fmt.Println("O tamanho da string é", len(s2))

	//char?? não existe char no GO. Ele apenas mapea o valor ascii do char e devolve um int32
	// var x char = 'a'
	char := 'a'
	fmt.Println("o tipo de char é", reflect.TypeOf(char))
}
