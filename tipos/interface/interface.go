package main

import (
	"fmt"
)
type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//Interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	//Programação estruturada: função trabalham com dados
	//OO: dados trabalham com funções
	var coisa imprimivel = pessoa{"Roberto", "Monteiro"}
	fmt.Println(coisa.toString())
	imprimir(coisa) //herança por composição

	coisa = produto{"Calça Jeans", 79.90} //Polimorfismo acontece aqui. Um hora é pessoa outra hora é produto
	fmt.Println(coisa.toString())
	imprimir(coisa) //herança por composição

	p1 := produto{"Calça Jeans", 99.90}
	imprimir(p1)
}


