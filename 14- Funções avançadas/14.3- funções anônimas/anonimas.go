package main

import "fmt"

func main() {

	func() {
		fmt.Println("Olá mundo!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando parâmetro")

	fmt.Println(retorno)
}
