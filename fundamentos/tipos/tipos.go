package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
1. int (int8...int64)
2. uint (uint8...uint64)
3. byte (alias for uint8)
4. rune (alias for uint32)
5. float32, float64
6. bool
7. string
8. complx64, complex128
*/
func main(){
//Numericos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))


	//Sem sinal(números naturais)... Unsigned uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1, reflect.TypeOf(i1))

	var i2 rune = 'a'//retorna o valor na tabela unicode
	fmt.Println("Rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

//Números Reais
	var x float32 = 49.90
	fmt.Println("O tipo de x é", reflect.TypeOf(x))

//Boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

//Strings
	s1 := "Anna Clara é Top"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1)) //O go não é OO, veja que len não é um método de classe nenhuma

	//String com multiplas linhas
	s2:= `Anna
	Clara
	é
	Top`
	fmt.Println("O tamanho da string é", len(s2)) //O valor se altera pois a funcao contas as Tabs

	//Golang tem char?
	char := 'a'
	//var x char = 'a' não existe
	fmt.Println("O tipo de char é", reflect.TypeOf(char)) //é rune que na verdade é int32

}