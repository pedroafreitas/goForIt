package main

import (
	"fmt"
	"reflect"
)

func main(){
	a1 := [...]int{1, 2, 3} //array
	s1 := []int{1, 2, 3} //slice - tamanho variável

	fmt.Printf("A1 %v Tipo %v\n",a1, reflect.TypeOf(a1))
	fmt.Printf("A1 %v Tipo %v\n",s1, reflect.TypeOf(s1))

	//Um slice é um pedaço de um vetor
	a2 := [5]int{2,3,4,5,6}
	s2 := a2[1:3]
	s3 := a2[3:]
	s4 := s3[:1]
	fmt.Println(a2, s2, s3, s4)

	s5 := make([]int, 10, 20)
	s6 := append(s5, 1, 2, 3)
	fmt.Println(s5, s6)

	s5[0] = 7
	fmt.Println(s5, s6) //vemos que os dois slices apontam para o mesmo vetor interno
	//Slice s2 e s3 apontam para o mesmo array. 
	//Slice aponta para um elemento de um array e 
	//a partir daquele elemento ele recebe os valores
	
}