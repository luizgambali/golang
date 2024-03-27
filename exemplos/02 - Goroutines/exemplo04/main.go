package main

import (
	"fmt"
	"time"
)

func primeiraFuncao() {
	for i := 0; i < 5; i++ {
		fmt.Println("Esta é uma mensagem da primeiraFuncao")
		time.Sleep(time.Second * 1)
	}
}

func segundaFuncao() {
	for i := 0; i < 5; i++ {
		fmt.Println("Esta é uma mensagem da segundaFuncao")
		time.Sleep(time.Second * 1)
	}
}

// a main é a rotina principal, roda na thread principal
func main() {

	//agora as funções vão rodar em paralelo
	go primeiraFuncao()
	go segundaFuncao()

	fmt.Scanln()
}
