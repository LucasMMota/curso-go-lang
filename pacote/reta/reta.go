package main

import "math"

// pacotes -> arquivos -> código
// no fim todos os arquivos viram todos uma estrutura de um <Pacote>
// quando compila junta tudo(arquivos) em uma estrutura só do pacote.
// pode ter vários arquivos de um mesmo pacote.

// Inicializando com letra maiúscula significa que será Público (visível dentro e fora do pacote)
// Inicializando com letra minúscula significa que será Privado (visível apenas dentro do pacote)

//Convenções
// Nome - gerará algo público
// nome ou _nome - gerará algo privado

// Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia calcula a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
