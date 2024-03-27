/*
	Código básico de uma aplicação Go.

	Toda aplicação Go tem como ponto de partida um main, tal qual o código abaixo.

	PRECISA ser package main
	PRECISA ter uma func main

	O código então será executado a partir das instruções dentro da função main

	para executar, basta abrir um terminal e digitar: go run <nome-do-arquivo.go>

*/

package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
