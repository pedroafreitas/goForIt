package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func Titulo(urls ...string) <-chan string {
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