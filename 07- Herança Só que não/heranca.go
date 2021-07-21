package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa      // basicamente essa é a "herança" em go
	curso       string
	instituicao string
}

func main() {
	fmt.Println("''Herança''")

	pessoa1 := pessoa{"Miguel", "Ártur", 1, 45}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Sistema de informação", "UFCA"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.sobrenome)
}
