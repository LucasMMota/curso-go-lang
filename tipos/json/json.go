package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"` // slice de strings
}

func main() {
	// struct to json
	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 1899.90,
		Tags:  []string{"Promoção", "Eletrônico"},
	}

	p1Json, _ := json.Marshal(p1) // transforma o struct para json, retorna []byte
	fmt.Println(string(p1Json))   // string([]byte) transforma em string

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"caneta", "preco": 8.9, "tags":["Papelaria", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
