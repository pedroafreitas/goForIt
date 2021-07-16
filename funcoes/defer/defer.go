package main

import (
	"fmt"
)

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto")
		return 5000 //defer é exec antes de cada return
	}
	fmt.Println("Valor baixo")
	return numero
	
}

func main(){
	fmt.Println(obterValorAprovado(5001))
	fmt.Println(obterValorAprovado(4999))
}