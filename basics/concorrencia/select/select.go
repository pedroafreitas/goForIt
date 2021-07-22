package main

import (
	"fmt"
	"time"
)

func oMaisRapido(url1, url2, url3 string, tempo float64) string {
	c1 := Titulo(url1)
	c2 := Titulo(url2)
	c3 := Titulo(url3)
	
	//Estrutura de controle especifica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(time.Millisecond * time.Duration(tempo)):
		return "Todos perderam\n"
	// default:
	// 	return "Sem respostas\n"
	}
}
func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
		1000,
	)

	fmt.Println(campeao)
}
