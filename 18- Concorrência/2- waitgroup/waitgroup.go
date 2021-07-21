package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //Dizendo que tem duas goroutines que fazem parte desse grupo de espera e ele precisa esperar terminar/ WaitGroup = Grupo de espera

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // Tira um do contador, dizendo que agora só tem que esperar mais 1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // Esperar que a contagem das goroutines chegue em zero
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
