package main

import (
	"fmt"
	"math"
)

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
	fmt.Println(notaParaConceito(11.0))
	fmt.Println(notaParaConceito(8.0))

	v := Vertex{3, 4}
	v = v.Scale(10)
	println(int(v.Abs()))
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { 
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
Neste exemplo nós precisamos retornar a struct Vertex se quisermos
passar os valores X e Y por cópia. Se tivessemos declarado:

func (v *Vertex) Scale(f float65) Vertex{}

teríamos uma função que recebe a referência dos valores reclarados
na main. Por isso, precisamos retornar a struct inteira e atribuir o
retorno a uma variável do mesmo tipo.
*/
func (v Vertex) Scale(f float64) Vertex{
	v.X = v.X * f
	v.Y = v.Y * f
	return Vertex{v.X,v.Y}
}
