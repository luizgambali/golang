package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	//go run main.go ip --host terra.com.br
	//go run main.go servidor --host google.com.br

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
