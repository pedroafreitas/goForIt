package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Terminou a rotina")
	wg.Done()
}

func main() {
	fmt.Println("Tutorial")
	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()
	fmt.Println("!!!Terminou o programa!!!")
}
