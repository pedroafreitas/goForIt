package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //operação bloqueante. Só prossegue quando alguém lê
	fmt.Println("Dado foi lido")
}

func main() {
	c := make(chan int) //canal sem buffer

	go rotina(c)

	fmt.Println(<-c) //operação bloqueante. Só recebo quando o dado está pronto
	fmt.Println("Foi lido")
	//fmt.Println(<-c) //deadlock#
	fmt.Println("FIM")
}
