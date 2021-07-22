package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	moedelo         string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	/*
		É mais comum no nível de interface não fazer
		alterações no que estamos acessando. Interfaces
		que causam efeitos colaterais são ruins.
		
		Em interfaces também busque criar interfaces com 
		métodos puros - de forma determinística.
	*/
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)

}
