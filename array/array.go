package main

import (
	"fmt"
)

func main(){
	//Homôgenea e estíatica
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	notas[2] = 3
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total/float64(len(notas))
	fmt.Printf("Média %f", media)

}