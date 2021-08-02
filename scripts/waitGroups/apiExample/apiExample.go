package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://tutorialedge.com",
	"https://twitter.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%v\n", err)
			}
			fmt.Fprintf(w, "%v\n", resp)
			wg.Done()
		}(url)
	}
	wg.Wait()
}
func main() {
	fmt.Println("Começando...")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
