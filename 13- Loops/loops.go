package main

import (
	"fmt"
	"time"
)

func main() {

	/* 	i := 0

	   	for i < 10 {
	   		i++
	   		fmt.Println("Incrementando i")
	   		time.Sleep(time.Second)
	   	}
	   	fmt.Println(i) */

	/* for j := 0; j < 10; j += 2 { //só fica vísivel dentro do escopo
		fmt.Println("Incrementando j de 2 em 2", j)
		time.Sleep(time.Second)
	} */

	nomes := [3]string{"Ana", "Amanda", "Vivian"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Andrel",
		"sobrenome": "Moura",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Não é possível utilizar range em struct

	//LOOP INFINITO
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
