package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {

    defer wg.Done()

    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    //fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wgAllScan sync.WaitGroup

    for i := 0; i <= 4; i++ {
        wgAllScan.Add(1)
        go worker(i, &wgAllScan)

        for j := 0; j <= 4; j++ {
            go worker(i, &wgAllScan)
        }
    }

    wgAllScan.Wait()
}