package main

import (
	"fmt"
)

// é executado primeiro no arquivo (antes do main inclusive)
func init() {
	fmt.Println("Inicializando")
}

func main() {
	fmt.Println("Main...")
}
