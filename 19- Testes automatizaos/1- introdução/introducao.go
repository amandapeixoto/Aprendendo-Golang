package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua João Pereira")
	fmt.Println(tipoEndereco)
}
