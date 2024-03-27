package main

import "fmt"

func main() {

	//se não especificar o buffer, para este caso aqui, ele vai dar dead lock.
	//Por que? Porque ele inicializa o buffer com zero. Ele vai colocar o valor
	//no canal, mas o tamanho é zero, então vai estourar!
	//Mesma situação se, na instrução abaixo, eu passar 3 valores para o canal:
	//como eu tenho um buffer de 2, se eu passar três vai estourar!

	//o canal com buffer, só vai bloquear quando chegar a sua capacidade; o canal
	//sem buffer vai bloquear a cada envio

	canal := make(chan string, 2)

	canal <- "Olá mundo!"

	mensagem := <-canal
	fmt.Println(mensagem)
}
