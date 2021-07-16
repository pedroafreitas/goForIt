package main

import (
	"fmt"
)

/*
Funcao pura gera sempre o mesmo resultado
dado os mesmos parametros e não gera nenhum
efeito colateral.
*/
func main(){
	f1()
	f2("Anna", "Clara")
	f3()
	s1 := f4("Anna", "Clara")
	println((s1))
	
	_, s2 := f5()

	println(s1, s2)
}
func f1() {
	fmt.Println("F1")
}

func f2(p1 , p2 string){
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string{
	return "F3"
}

func f4 (p1, p2 string) string{
	return fmt.Sprintf("F4 %s %s", p1, p2) //retorna string formatada
}

//Retornando multiplos valores
func f5() (string, string){
	return "Anna", "é Top"
}