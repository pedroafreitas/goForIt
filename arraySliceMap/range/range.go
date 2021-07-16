package main

import (
	"fmt"
)
func main() {
	//O compilador vai contar o tamanho do vetor
	numero := [...]int{1, 2, 3, 4, 5}

	for i, numero := range numero {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, numero := range numero {
		fmt.Println(numero)
	}

	
}