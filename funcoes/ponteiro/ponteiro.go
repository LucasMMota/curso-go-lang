package main

import (
	"fmt"
)

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// * é usado aqui para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	inc2(&n) // passando por referência
	fmt.Println(n)
}
