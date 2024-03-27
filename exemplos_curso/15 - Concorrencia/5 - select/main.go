package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		//da forma abaixo, ele vai receber a mensagem do canal 1,
		//mas a execução vai ficar bloqueada esperando a mensagem do cancal 2
		// mensagemCanal1 := <-canal1
		// mensagemCanal2 := <-canal2

		// fmt.Println(mensagemCanal1)
		// fmt.Println(mensagemCanal2)

		//desta forma, ele não fica esperando, e executa a instrução conforme
		//o canal que está recebendo a mensagem
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
