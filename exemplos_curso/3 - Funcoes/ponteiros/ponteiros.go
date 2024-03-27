package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}
func main() {

	numero := 10

	fmt.Println("Valor original: ", numero)
	numeroInvertido := inverterSinal(numero)
	fmt.Println("Valor invertido: ", numeroInvertido)
	fmt.Println("Valor original: ", numero)

	novoNumero := 40

	fmt.Println("Valor original: ", novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println("Valor invertido: ", novoNumero)
}
