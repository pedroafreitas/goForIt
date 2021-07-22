package main

import "fmt"

//prefira composição no lugar de herança
//vamos fazer uma pseudoherança usando composição

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	//campo anonimo: vai "herdar" tudo de carro
	carro
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Println(f)
}
