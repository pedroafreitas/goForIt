package main

import (
	"fmt"
	"time"
)

func main() {

	//Operadores básicos
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2

	fmt.Println(i)

	x, y := 1, 2 //Inicializando duas variáveis
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)

	//Operadores Relacionais
	fmt.Println("==", "a" == "n")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 9)
	fmt.Println(">=", 4 >= 5)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println(d1)
	fmt.Println("Data", d1 == d2)
	fmt.Println("Datas", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoa", p1 == p2)

	//Operadores Lógicos
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete %t, \"Saudável\": %t\n", tv50, tv32, sorvete, !sorvete)



	//Operadores Unários
	add := 1
	add++
	fmt.Println(add)

	add--
	fmt.Println(add)
	
	//Operadores "Ternários"
	fmt.Println(obterResultado(6.5))


}

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	comprarSorverte := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorverte
}

func obterResultado(nota float64) string{
	if nota >= 6{
		return "Aprovado"
	}
	return "Reprovado"
	//Não existe
	//return nota >= 6 ? "Aprovado" : "Reprovado"
}