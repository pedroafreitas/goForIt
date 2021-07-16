package main

import (
	"fmt"
	"runtime/debug"
)

func f3() {
	debug.PrintStack()
}

func f2(){
	f3()
}

func f1() {
	f2()
}

func main(){
	f1()

	x, y := troca(2,3)
	fmt.Println(x, y)
}

func troca(p1,p2 int) (segundo, primeiro int){
	segundo = p2
	primeiro = p1
	return //retorno limpo, retorna as variaveis de retorno sem colocar os nomes
}