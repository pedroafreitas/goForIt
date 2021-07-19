package main

import (
	"time"
	"fmt"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++{
		for primo := inicio; ; primo++{
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Microsecond * 10000)
				break
			}
		}
	}
	close(c) //se não fecha o canal, vai dar deadlock.  O for na main que gera o deadlock por que o for fica esperando receber alguma coisa.
}

func main() {
	c := make(chan int, 3000)
	go primos(cap(c), c)

	for primo := range c {
		fmt.Printf("%d ", primo)
	}
	fmt.Println("FIM")
}
