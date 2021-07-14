package main

import (
	"fmt"
	r "reflect" //adiciona um apelido para o pacote
	"strconv"
)
func main(){
	x := 2.4
	y := 2
	//fmt.Println(x/y) dividindo tipos diferentes
	fmt.Println(x / float64(2))
	fmt.Println(int(x)/ y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado!!!
	fmt.Println("Teste " + string(97))

	//int para string
	fmt.Println("Teste 2 " + strconv.Itoa(123))
	
	//string para int
	num, _ := strconv.Atoi("123") //uma função em Go pode retornar dois valores
	//"_" ignora a variável
	fmt.Println("O tipo de num é", r.TypeOf(num))

	b, _ := strconv.ParseBool("123") //os unicos valores veradeiros são true e 1

	if b {
		fmt.Println("Bool e verdadeiro")
	} else {
		fmt.Println("Bool e falso")
	}

	b, _ = strconv.ParseBool("True")

	if b {
		fmt.Println("Bool é verdadeiro")
	} else {
		fmt.Println("Bool é falso")
	}
}

