package main

import (
	"fmt"
	"time"
)

func main() {

	//canal é uma variável do tipo channel que vai receber um int
	//ela vai permitir a troca de mensagens entre as goroutines
	var canal = make(chan string)

	//funcao anonima 1
	go func() {
		defer close(canal) //ao terminar a goroutine, o canal é fechado, evitando deadlocks

		for i := 0; i < 5; i++ {
			canal <- fmt.Sprintf("%v - %v", "primeira funcao: ", i)
			time.Sleep(time.Second * 1)
		}
	}()

	//funcao anonima 2
	go func() {
		defer close(canal) //ao terminar a goroutine, o canal é fechado, evitando deadlocks

		for i := 0; i < 5; i++ {
			canal <- fmt.Sprintf("%v - %v", "segunda funcao: ", i)
			time.Sleep(time.Second * 1)
		}
	}()

	for x := range canal {
		println(x)
	}

}
