package main

import "fmt"

func main() {
	canal := make(chan string, 2) //Usando buffer. Está dizendo que tem a capacidade de 2 canais(Bloqueia quando atingir a capacidade máxima)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
