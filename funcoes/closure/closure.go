package main

import "fmt"

func main(){
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX() 
	/*
	A função "respeita as origens". A função consegue
	saber o contexto léxico em que ela foi definida.
	
	A função não traz apenas seus comandos, mas também
	seu contexto.
	*/
}

func closure() func(){
	x := 10
	var funcao = func(){
		fmt.Println(x)
	}
	return funcao
}