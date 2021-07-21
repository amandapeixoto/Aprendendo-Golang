package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)

	/* for {
		mensagem, aberto := <-canal
		if !aberto { // avalia se ainda está aberto, se não tiver aberto sai do loop
			break
		}
		fmt.Println(mensagem)
	} */

	for mensagem := range canal { // Faz a mesma coisa que o for de cima
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // essa linha envia o valor para <- canal
		time.Sleep(time.Second)
	}

	close(canal) // fechar canal de comunicação após as 5x
}
