package main

import (
	"fmt"
)

func imprimir(valor int){
	fmt.Println(valor)
}

func somar(a int, b int) int{
	return a + b
}

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorverte := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorverte
}