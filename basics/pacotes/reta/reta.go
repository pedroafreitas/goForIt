package main

import "math"

/*
Inicializandoo com letra maiúscula é PÚBLICO
Público: com visibilidade dentro e fora do pacote

Não existe visibilidade privada dentro do arquivo, mas
dentro do pacote

Inicializando com minúscula é PRIVADO
Privado: com visibilidade APENAS DENTRO do pacote

Exemplos:
1. Ponto - gerará algo público
2. ponto ou _Ponto - gerará algo privado
*/

//Ponto representa uma coordenada do plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//Distância calcula a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

