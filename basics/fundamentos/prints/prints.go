package main

import "fmt"

func main() {

	x := 3.141516

	xString := fmt.Sprint(x) //Retorna String

	//O "+" serve para concatenar Strings
	fmt.Println("O valor de x é " + xString)
	//"," concatena ...
	fmt.Println("O valor de x é", x, "!")

	fmt.Printf("O valor de x é %.2f\n", x)

	a := 1
	b := 1.999
	c := "opa"
	d := false

	fmt.Printf("%d %f %.1f %s %t\n", a, b, b, c, d)


}
