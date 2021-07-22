package main

import "fmt"

type item struct {
	produtoID int
	qtd int
	preco float64
}

type pedido struct {
	userID int
	itens []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := pedido {
		userID: 1,
		itens: []item{
			{1, 2, 12.10}, //essa forma reduzida por de problematica
			{2, 1, 100.00},
			{3, 10, 2.99},
		},
	}

	fmt.Printf("Valor total: R$ %.2f\n", pedido.valorTotal())
}