package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

//Método = está obrigatoriamente ligado a uma estrutura, não é solto como as funções
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Amanda", 25}
	usuario1.salvar()

	usuario2 := usuario{"Romão", 33}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
