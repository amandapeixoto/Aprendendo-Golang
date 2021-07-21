package main

import "fmt"

func main() {
	fmt.Println("Função main sendo executada")
}

//Pode ter uma função init por arquivo e não por pacote

func init() {
	fmt.Println("Função init sendo executada")
}
