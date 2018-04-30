package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	defer close(c) // fecha o canal, senão no for vai tentar ler e va idar deadlock pois o canal estará vazio

	inicio := 2
	for i := 0; i < n; i++ { // pega os n primeiros primos
		for primo := inicio; ; primo++ { // testa cada primo a partir do numero anterior(inicio)
			if isPrime(primo) {
				c <- primo // se é primo, poe no buffer
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
}

func main() {
	c := make(chan int, 30) // canal com buffer de tamanho 30
	go primos(10, c)
	for primo := range c { // le cada item do canal e imprime
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
