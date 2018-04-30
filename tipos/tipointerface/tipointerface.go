package main

import (
	"fmt"
	"reflect"
)

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{} // interface vazia

	var coisa2 dinamico = "Opa" // defino coisa2 do tipo dinamico e implicitamente defino como string
	fmt.Println(reflect.TypeOf(coisa2), coisa2)

	coisa2 = true // agora atribuo um bool e coisa2 vira boolean
	fmt.Println(reflect.TypeOf(coisa2), coisa2)

	coisa2 = curso{"Curso GO Lang"}
	fmt.Println(reflect.TypeOf(coisa2), coisa2)

	// quando a var é do tipo interface{} pode-se atribuir vários tipos a ela durante a execucao
}
