package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) { // vai inferir o tipo do dado
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Não sei..."
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Uooopaaa!!!"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
