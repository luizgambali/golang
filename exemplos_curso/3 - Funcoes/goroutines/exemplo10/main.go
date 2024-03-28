package main

import "fmt"

func funcaoGenerica() chan string {
	var canal = make(chan string)

	go func() {
		defer close(canal)

		for i := 0; i < 5; i++ {
			canal <- fmt.Sprintf("%v - %v", "funcao generica", i)
		}
	}()

	return canal
}

func main() {
	var canal = funcaoGenerica()

	for x := range canal {
		println(x)
	}
}
