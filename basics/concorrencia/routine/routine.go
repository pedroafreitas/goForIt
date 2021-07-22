package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s [%d]\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Por que você não comigo?", 3)
	// fale("João", "Só posso falar depois de você", 1)

		go fale("Maria", "Ei...", 500)
		go fale("João", "Opa!", 500)

		time.Sleep(time.Second * 10)
		fmt.Println("Fim!")

		go fale("Maria", "Entendi!!!",10)
		fale("João", "Parabéns", 5)
}
