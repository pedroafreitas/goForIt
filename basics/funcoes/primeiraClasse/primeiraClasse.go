package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main(){
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(100,10))

	result := exec(multi, 4, 4)
	fmt.Println(result)
}

func multi(a, b int) int{
	return a * b
}

func exec(funcao func(int, int) int , p1, p2 int) int{
	return funcao(p1, p2)
}