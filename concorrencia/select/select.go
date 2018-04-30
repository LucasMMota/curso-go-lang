package main

import (
	"fmt"
	"time"

	"github.com/lucasmmota/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1: // se o valor do canal chegar
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(100 * time.Millisecond): // se depois de 1 segundo nenhum chegar
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)

	fmt.Println(campeao)
}
