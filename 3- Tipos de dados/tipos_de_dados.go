package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	/* 	uint não suporta sinal, exemplo = -100 */
	var numero2 uint = 100
	fmt.Println(numero2)

	/* 	alias: rune(apelido para int32)
	alias: byte(apelido para int8) */
	var numero3 rune = 98765
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)
	// Fim números inteiros

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12334546.46
	fmt.Println(numeroReal2)
	// Fim números reais

	var caracteres string = "Texto, números= 1234"
	fmt.Println(caracteres)

	caracteres2 := "Texto"
	fmt.Println(caracteres2)
	// Fim strings

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
