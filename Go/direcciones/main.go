package main

import "fmt"

func main() {
	numerosLista := numeros{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numerosLista.printearNumerosParImpar()
}

type numeros []int

func (n numeros) printearNumerosParImpar() {

	for _, numero := range n {
		if numero%2 == 0 {
			fmt.Println(numero, "es par")
		} else {
			fmt.Println(numero, "es impar")
		}
	}
}
