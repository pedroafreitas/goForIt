package main

import (
	"fmt"
	_ "math" //Importa e não usa
)

func main() {
	//const <nome> <tipo>
	const PI float64 = 3.1415
	var raio = 3.7 //Valor inferido (float64)

	//:= declara e inicializa
	area := 2 * PI * raio

	//Sempre que declaramos uma variável devemos usá-la
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
