package main

import (
	"fmt"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3 // aqui encheu o canal
	ch <- 4 // canal cheio: aguarda liberar espaço para continuar a execução
	ch <- 5
	fmt.Println("Executou!")
	ch <- 6
}
func main() {
	ch := make(chan int, 3) // canal com buffer de tamanho 3
	go rotina(ch)           // executa enquanto o código estiver executando

	fmt.Println(<-ch) // imprime o primeio valor do canal
	// não há problema em não ler todos os dados
}
