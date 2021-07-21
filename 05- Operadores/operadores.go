package main

import "fmt"

func main() {
	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	//Os cálculos só são realizados se as variáveis tiverem o mesmo tipo de dados, ex: int 8

	// FIM DOS ARITMÉTICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// FIM DAS ATRIBUIÇÕES

	// RELACIONAIS (SEMPRE RETORNAM TRUE OU FALSE)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// FIM DOS RELACIONAIS

	// LÓGICOS
	fmt.Println("_______________")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro) // A exclamação inverte o resultado
	fmt.Println(!falso)

	// FIM DOS LÓGICOS

	// UNÁRIOS
	numero := 10
	numero++
	fmt.Println(numero)

	numero1 := 10
	numero1++
	numero1 += 15 // numero = numero + 15
	fmt.Println(numero1)

	numero2 := 10
	numero2--
	numero2 *= 3 // numero2 = numero2 * 3 (Pode utilizar com todos os operadores lógicos)
	fmt.Println(numero2)

	// FIM UNÁRIOS

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menos que 5"
	}
	fmt.Println(texto)
}
