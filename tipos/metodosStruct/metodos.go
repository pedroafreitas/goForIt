package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

//se não tiver o ponteiro vou estar alterando a varivel interna ao escopo.
//A mudança só ocorre na main por causa do ponteiro
func (p *pessoa) setNomecompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"Pedro", "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomecompleto("Anna Vieira")
	
	fmt.Println(p1.getNomeCompleto())
}
