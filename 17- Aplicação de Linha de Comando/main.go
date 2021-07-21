package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
	/* erro := aplicacao.Run(os.Args) //Também pode ser assim
	if erro != nil {
		log.Fatal(erro)
	} */
}
