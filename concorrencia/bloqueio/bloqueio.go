package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("Só executa depois que for lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operação bloqueando
	fmt.Println("Foi lido")
	//fmt.Println(<-c) // deadlok
	fmt.Println("Fim")
}
