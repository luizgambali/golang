package main

import (
	"fmt"
	"time"
)

func mostrarMensagem() {
	fmt.Println("Agora essa mensagem será exibida, porque o programa principal está esperando a goroutine terminar")
}

// a main é a rotina principal, roda na thread principal
func main() {

	// executa a função mostrarMensagem como uma goroutine
	go mostrarMensagem()

	// exibe a mensagem antes da rotina
	fmt.Println("Esta mensagem será exibida antes da rotina")

	// espera um tempinho, para a goroutine rodar
	time.Sleep((time.Second * 2))
}
