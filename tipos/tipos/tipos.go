package main

import "fmt"

type nota float64 //cria um alias

//passa como receiver
func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.00, 6.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	case n.entre(0.0, 3.0):
		return "E"
	default:
		return "Inválido"
	}
}

func main() {
	/*
	Aqui nós estamos criando um método "entre" para que
	ele não precise ser reutilizado posteriormente por outras
	funções. 
	*/
	fmt.Println(notaParaConceito(3.3))

}
