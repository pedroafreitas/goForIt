package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4 //aqui vai ficar bloquado, mas as vezes executa o print abaixo.
	fmt.Println("Executou")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch) //deadlock é quando não tem mais ninguém para enviar dado e vc fica esperando receber

}
