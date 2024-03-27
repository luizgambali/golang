package main

import "fmt"

func LoopWhile() {

	fmt.Println(" *** WHILE ***")
	var soma = 5

	for soma > 0 {
		fmt.Println("Valor: ", soma)
		soma--
	}
}

func LoopNormal() {

	fmt.Println(" *** LOOP NORMAL ***")
	for i := 0; i < 10; i++ {
		fmt.Println("Estrutura for: iteracao ", i)
	}

}

func LoopForEach() {
	dias := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"}

	for indice, valor := range dias {
		fmt.Println("Posicao ", indice, "Valor", valor)
	}
}

func LoopDoWhile() {
	fmt.Println(" *** DO WHILE ***")
	var soma = 5

	for {
		fmt.Println("Valor: ", soma)
		soma--

		if soma == 0 {
			break
		}

	}
	fmt.Println("Saiu do laço ok!")
}

func main() {

	//go não possui os laços while e do while, somente for.
	//porém há formas de replicar isso

	LoopWhile()

	LoopNormal()

	LoopForEach()

	LoopDoWhile()
}
