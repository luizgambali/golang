package main

import (
	"fmt"
	"time"
)

func main() {
	//cria um canal (que só recebe dados) dentro de escrever, e retorna para a variavel canal
	canal := escrever("Olá mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

//a rotina escrever pode ser chamada normalmente, sem ser uma go routine
//a goroutine fica encapsulada dentro da função escrever, e a função
//escrever vai retornar um canal de comunicação
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)

		}
	}()

	return canal
}
