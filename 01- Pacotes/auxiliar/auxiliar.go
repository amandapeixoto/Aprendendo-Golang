package auxiliar

import (
	"fmt"
)

/* Se uma função começar com letra minúscula ela só será visível no pacote que ela está */
/* Se começar com letra maiúscula pode ser exportada para funções de outros pacotes */
//Escrever: registra uma mensagem na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
