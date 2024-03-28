package main

import "fmt"

func mostrarMensagem() {
	fmt.Println("Esta rotina será executada como uma goroutine")
}

// a main é a rotina principal, roda na thread principal
func main() {

	// executa a função mostrarMensagem como uma goroutine
	go mostrarMensagem()

	// exibe a mensagem antes da rotina
	fmt.Println("Esta mensagem será exibida antes da rotina")

	// a mensagem da goroutine não foi exibida, pois o programa principal terminou antes de ela ser executada
}
