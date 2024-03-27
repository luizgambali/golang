package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Concorrência é diferente de paralelismo

		Em paralelismo, cada código é executado em uma thread única, e são executadas ao mesmo tempo
		em paralelo

		Concorrência também distribui a execução, mas ficaria como se fosse uma mesma thread, executa
		um pouco uma coisa, depois executa um pouco outra

		São conceitos bem similares, mas tem diferença

		Colocando o comando "go" em frente a função, ele não espera o término da execução daquela função,
		para que o código abaixo seja executado

	*/

	go escrever("Hello world!")
	escrever("Bye world!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
