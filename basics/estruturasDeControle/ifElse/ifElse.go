package main

import ( 
	"math/rand"
	"time" 
	"fmt"
)

func main() {
	imprimirResultado(6.5)
	fmt.Println(notaParaConceito(8.3))

	//If com init
	if i := aleatorio(); i > 5 { //variável local
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Perdeu")
	}
}
func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func notaParaConceito(n float64) string {
	//Nesse caso seria melhor usar switch
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func aleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}