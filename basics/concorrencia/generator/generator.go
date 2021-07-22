package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//GOOGLE I/O 2021 - GO CONCURRENCY PATTERNS

//<- chan - canal somente-leitura

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
		/*Essa função recebe um parametro. Por isso eu passo url.
		Eu estou criando uma função e ao mesmo tempo invocando-a.
		Cada vez que o laço executa eu chamo uma nova go routina
		para que ele execute a função concorrentemente.*/
	}
	return c
}

func main() {
	/*
	Esse padrão já encapsula as routines dentro da função.
	Não me preocupo com buffer, a criação do canal ou qualquer
	outra coisa. A única tarefa que eu faço é chamar o canal.

	É muito mais rápido eu fazer isso com concorrência do que
	de forma serial.
	*/
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
