package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA(REVEZAMENTO, UTILIZANDO O MESMO NÚCLEO) != PARALELISMO (TAREFAS EXECUTADAS AO MESMO TEMPO, SE O PC TIVER MAIS DE UM NÚCLEO)
	go escrever("Olá Mundo!") //goroutines (não esperar terminar a execução desta linha independente do loop, e já passar para a próxima.)
	escrever("Programando em Go!")

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
