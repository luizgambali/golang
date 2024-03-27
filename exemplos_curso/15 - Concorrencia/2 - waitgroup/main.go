package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	//indica q serão duas rotinas rodando em concorrencia
	//vai ser usado com o Wait() no final da função
	waitGroup.Add(2)

	//função anonima
	go func() {
		escrever("Olá Mundo")
		waitGroup.Done() //retira 1 do waitGroup definido lá no Add
	}()

	//função anonima
	go func() {
		escrever("Tchau Mundo")
		waitGroup.Done() //retira 1 do waitGroup definido lá no Add

	}()

	//espera a execução de toda a lista do waitGroup antes de terminar o programa/rotina
	//como lá em cima foram adicionadas 2, ele vai esperar o fim da execução dessas duas,
	//que terminam com o Done()
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
