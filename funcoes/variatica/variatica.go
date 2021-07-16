package main

import "fmt"


func media(numeros ...float64) float64{
	total := 0.0
	for _, num := range numeros{
		total += num
	}
	return total / float64(len(numeros))
}

func imprimirAprovados(aprovados ...string) {
	fmt.Println("\nLista de aprovados:")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}
func main() {
	fmt.Printf("Média %.2f\n", media(1,2,3,4,5,6,7,8,9))
	
	aprovados := []string{"Anna", "Pedro", "João", "Carlos"}
	imprimirAprovados(aprovados...)

}