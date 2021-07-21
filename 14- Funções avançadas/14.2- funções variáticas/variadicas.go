package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 5, 8, 9, 10, 15, 20)
	fmt.Println(totalDaSoma)

	escrever("Olá", 4, 6, 7)
}
