package main

import (
	"fmt"
)

type produto struct {
	//Um agrupador de dado
	nome     string
	preco    float64
	desconto float64
}

//receiver da função
//"quem será o dono dessa função?"

//Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	//var produto1 produto
	produto1 := produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1)

	produto2 := produto{"Notebook", 2000.00, 0.10}
	produto2.preco -= 100
	fmt.Println(produto2.precoComDesconto())

	
}
