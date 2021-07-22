package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(0))
	fmt.Println(notaParaConceito(-1))

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}

	fmt.Println(tipo(2.3))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(time.Now()))
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "interiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "Tipo desconhecido"
	}
}

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch{
	case nota<0 || nota > 10:
		return "Nota inválida"
	case nota>= 8:
		return "A"
	case nota>= 8:
		return "B"
	case nota>= 5:
		return "C"
	case nota>= 3:
		return "D"
	default:
		return "E"
	}
	// switch nota{ //Se executar um case, a linguagem sai do switch
	// case 10:
	// 	fallthrough
	// case 9:
	// 	return "A"
	// case 8, 7:
	// 	return "B"
	// case 6, 5, 4:
	// 	return "C"
	// case 3,2,1,0:
	// 	return "E"
	// default:
	// 	return "Nota inválida"
	// }
}