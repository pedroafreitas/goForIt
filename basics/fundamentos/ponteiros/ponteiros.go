package main

import (
	"fmt"
)
func main() {
	i := 1

	//Go não tem aritmética de ponteiros
	var p *int = nil

	p = &i //atribuindo o endereço de i à p

	*p++ //deferenciando o ponteiro para pegar o valor
	i++
	*p++
	
	fmt.Println(i)
	fmt.Println(p, *p)

}