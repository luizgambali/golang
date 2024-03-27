package main

import "fmt"

/*
	tipos de dados em Go

	Tipos primitivos: numeros, string, booleanos

		Numeros: inteiros ou com casas decimais

			int, int8, int16, int32 e int64
			uint, uint8, uint16, uint32, uint64 (unsigned int)

			float32, float64

			atenção ao tipo int: se usado tipo int, ele assumirá o valor
			de acordo com a arquitetura do processador. X32 ou X64

		Strings:

			string

		Booleano:

			true ou false

	Tipos complexos: são tipos criados pelo usuario, ou tipos especiais da linguagem

		nil, func, struct, array, slice, channel

*/

func main() {

	//declaração explicita do tipo
	var nome string = "João"
	var idade int8 = 30
	var salario float32 = 1890.55
	var possuiCarro bool = true

	fmt.Println(nome, idade, salario, possuiCarro)

	//declaração implicita do tipo

	endereco := "Rua das Flores, 15"
	principal := false

	fmt.Println(endereco, principal)
}
