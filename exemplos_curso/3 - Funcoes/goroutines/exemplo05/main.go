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

func terceiraFuncao() {
	for i := 0; i < 100; i++ {
		fmt.Println("Mensagem da terceira funcao")
	}
}

// a main é a rotina principal, roda na thread principal
func main() {

	//agora as funções vão rodar em paralelo
	//note que, como a terceira função não tem uma interrupção, ela	vai ser executada até o final

	go primeiraFuncao()
	go segundaFuncao()
	go terceiraFuncao()

	fmt.Scanln()
}
