package main

import (
	"fmt"
	"time"
)
func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i :=0; i <= 20; i +=20 {
		fmt.Printf("%d\n", i)
	}

	for i :=1; i <=10; i++  {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar")
		}
	}

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}