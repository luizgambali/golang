package main

import (
	"fmt"
	"time"
)

func primeiraFuncao() {
	for i := 0; i < 5; i++ {
		fmt.Println("Esta é uma mensagem da primeiraFuncao")
	}
}

func segundaFuncao() {
	for i := 0; i < 5; i++ {
		fmt.Println("Esta é uma mensagem da segundaFuncao")
	}
}

// a main é a rotina principal, roda na thread principal
func main() {

	// as mensagens da segunda função vão aparecer antes da primeira,
	// porque não existe um controle de sincronização entre as goroutines
	go primeiraFuncao()
	go segundaFuncao()

	time.Sleep(time.Second * 2)
}
