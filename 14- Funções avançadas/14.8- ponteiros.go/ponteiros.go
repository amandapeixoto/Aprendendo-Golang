package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1 //PASSANDO UMA CÓPIA
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1 //PASSANDO UMA REFERÊNCIA, QUALQUER ALTERAÇÃO FEITA VAI IMPACTAR A VARIÁVEL DA FUNÇÃO TAMBÉM (PARA GARANTIR QUE A ALTERAÇÃO SEJA FEITA NO CÓDIGO INTEIRO)
}

func main() {
	numero := 15
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
