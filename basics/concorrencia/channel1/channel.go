package main

import "fmt"

func main() {
	ch := make(chan int, 1) //criando um canal de inteiros

	ch <- 1 //enviando dados para o canal(escrita)
	<-ch    //recebendo dados do canal (leitura)

	ch <- 2 //enviando informação
	fmt.Println(<-ch)

	/*
	A ideia básica do canal é ter uma estrutura que
	pode ter dados para leitura e escrita. Vamos usar
	isso juntos com as routines para leiamos os
	dados da routina. 

	As funções das routines serão executadas assicronamentes.
	
	A partir do canal, temos sincronia. A partir do canal
	podemos esperar os dados que estão rodando nas routines.
	*/
}
