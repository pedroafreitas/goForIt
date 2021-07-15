package main

import (
	"fmt"
)

func main(){
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	//Capacidade é diferente de tamanho
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 0,0,0,0,0,0,0,0,0,0,0,0,0,0)
	fmt.Println(s, len(s), cap(s))

	for i := 0; i < len(s)/2; i++{
		s[i] = i+1
	}

	fmt.Println(s, len(s), cap(s))

	
}