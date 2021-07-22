package main

import (
	"fmt"
)

/*
Multiplexar é pegar duas fontes de dados diferentes
em apenas uma. Pegamos dois canais e juntamos em um.
*/

//Sempre que chegar dados no canal de origem, encaminho para o canal destino.
func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem //origem é canal de somente leitura
	}
}

//multiplexar - mistura mensagens num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
