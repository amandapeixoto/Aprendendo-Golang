package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 { // função de parada
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func fibonacci1(posicao1 uint) uint {
	if posicao1 <= 1 {
		return posicao1
	}

	return fibonacci1(posicao1-2) + fibonacci(posicao1-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(9)
	fmt.Println(fibonacci(posicao))

	posicao1 := uint(15)

	for i := uint(1); i <= posicao1; i++ {
		fmt.Println(fibonacci1(i))
	}
}
