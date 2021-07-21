package main

import "fmt"

/* func recuperarExecucao() {
	fmt.Println("Tentativa de recuperar a execução")
} Printa a frase na tela mas não resolve o problema do panic*/

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente igual a 6") //Antes do panic ele exibe o defer
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 5))
	fmt.Println("Pós execução")
}
