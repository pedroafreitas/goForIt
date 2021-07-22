package main

import (
	"fmt"
)

func inc1(i int) {
	i++
}

func inc2(i *int) {
	//Ponteiro é representado por *
	//* é usado para acessar o valor do ponteiro
	*i++
}

func main(){
	n := 1
	
	inc1(n) //passando por copia
	fmt.Println(n)

	inc2(&n) //passando por referencia
	fmt.Println(n)
}