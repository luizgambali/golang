package main

import "fmt"

func RecuperarErro() {
	if r := recover(); r != nil {
		println("Recuperado com sucesso!")
	}
}

func GerarErroRecuperar() {
	defer RecuperarErro()

	mensagem := "Olá, mundo!"

	fmt.Println(mensagem)

	panic("Deu um erro aqui!")
}

func GerarErroSemRecuperar() {
	panic("Agora não vai ter jeito!")
}

func main() {

	GerarErroRecuperar()

	fmt.Println("Continuando a execução do programa...")

	GerarErroSemRecuperar()
}
