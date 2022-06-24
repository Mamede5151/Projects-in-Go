package main

import (
	"Linha-de-Comando/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("POnto de Partida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
