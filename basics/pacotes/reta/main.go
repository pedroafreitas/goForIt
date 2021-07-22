package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	fmt.Println(catetos(p1, p2)) //disponível dentro do pacote
	fmt.Println(Distancia(p1, p2)) //disponível para fora do pacote
}
