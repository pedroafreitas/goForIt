package main

import (
	"fmt"
)

func main () {
	aprovados := make(map[int]string)

	aprovados[12345678] = "Anna"
	aprovados[12345670] = "Pedro"
	aprovados[12345600] = "João"
	fmt.Println(aprovados)

	for cod, nome := range aprovados {
		fmt.Printf("%s (COD: %d)\n", nome, cod)
	}

	fmt.Println(aprovados[456]) //imprime linha vazia
	fmt.Println(aprovados[12345600])
	delete(aprovados, 12345678)

	funcESalario := map[string]float64{
		"Pedro": 2000.00,
		"Anna": 400.00,
		"João": 1000.00,
	}

	funcESalario["Elane"] = 5000.00
	delete(funcESalario, "NãoExiste")

	for nome, salario := range funcESalario{
		fmt.Println(nome,salario)
	}

	//Map Aninhado
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 3213.00,
			"Guga": 4343.00,
		},
		"A": {
			"Anna": 30000.00,
			"Amanda": 100.00,
		},
		"P": {
			"Pedro": 1200.00,
		},
	}

	fmt.Println(funcPorLetra)

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
		for nome, valor := range funcPorLetra[letra] {
			fmt.Println(nome, valor)
		}
	}
}