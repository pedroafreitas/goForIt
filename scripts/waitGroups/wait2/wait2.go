package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	defer wg.Done()
}
func worker2(id int, wg *sync.WaitGroup) {
	errs, _ := errgroup.WithContext(context.Background())
	
	errs.Go(func() error {
		for i := 0; i <= 1; i++ {
			fmt.Printf("Processando erro %d\n", i)
			time.Sleep(time.Second)
		}
		return nil
	})
	errs.Wait()
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	var wg1 sync.WaitGroup
	for i := 1; i <= 3; i++ { //3x
		wg.Add(1)
		go worker(i, &wg)
		for i := 3; i <= 4; i++ { //2x
			wg1.Add(1)
			go worker2(i, &wg1)
		}
	}
	wg1.Wait()
	wg.Wait()
}
