package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre as goroutines
// channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	a, b := <-c, <-c // só executa depois que pegar os dois valores do canal c
	fmt.Println(a, b)

	fmt.Println(<-c)
}
