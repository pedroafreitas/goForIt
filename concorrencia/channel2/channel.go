package main

import (
	"fmt"
	"time"
)

//Channel - forma de comunicação entre as routines
//channel também é um tipo - cidadão de primeira linha
func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	
	a, b := <-c, <-c //recebendo dados do canal
	fmt.Println(a, b)
	fmt.Println("B")

	fmt.Println(<-c)
}

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}
