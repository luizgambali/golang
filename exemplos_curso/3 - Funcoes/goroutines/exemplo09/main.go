package main

import (
	"fmt"
	"time"
)

func primeiraFuncao(canal chan string) {
	defer close(canal) //ao terminar a goroutine, o canal é fechado, evitando deadlocks

	for i := 0; i < 5; i++ {
		canal <- fmt.Sprintf("%v - %v", "primeira funcao: ", i)
		time.Sleep(time.Second * 1)
	}
}

func segundaFuncao(canal chan string) {
	defer close(canal) //ao terminar a goroutine, o canal é fechado, evitando deadlocks

	for i := 0; i < 5; i++ {
		canal <- fmt.Sprintf("%v - %v", "segunda funcao: ", i)
		time.Sleep(time.Second * 1)
	}
}

// mesma coisa que o exemplo08, mas com funções separadas

func main() {

	var canal = make(chan string)

	go primeiraFuncao(canal)
	go segundaFuncao(canal)

	for x := range canal {
		println(x)
	}

}
