package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//junta os dois canais em um só
	//escrever retorna um canal... então passa os dois para a rotina
	canal := multiplexar(escrever("Ola Mundo"), escrever("Programando em Go!"))

	//mostra as mensagens do canal
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalentrada1, canalentrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalentrada1:
				canalSaida <- mensagem
			case mensagem := <-canalentrada2:
				canalSaida <- mensagem

			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			//time.Duration(rand.Intn(2000)) -> faz uma duração aleatória, entre 0 e 2000 ms
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))

		}
	}()

	return canal
}
