package main

import "fmt"

func RetornaElemento(lista []string, posicao int) string {

	if posicao > 0 && (posicao-1) < len(lista) {
		return lista[posicao-1]
	} else {
		return ""
	}

}

func main() {

	var dia []string

	//adiciona elementos ao slice
	dia = append(dia, "Segunda")
	dia = append(dia, "Terça")
	dia = append(dia, "Quarta")
	dia = append(dia, "Quinta")
	dia = append(dia, "Sexta")

	//equivalente ao "foreach"
	for _, valor := range dia {
		fmt.Println(valor)
	}

	fmt.Println("Elemento na posicao 1: ", RetornaElemento(dia, 1))
	fmt.Println("Elemento na posicao 5: ", RetornaElemento(dia, 5))

	mes := []string{"Janeiro", "Fevereiro", "Março", "Abril"}
	fmt.Println(mes, len(mes))
	mes = append(mes, "Maio")
	fmt.Println(mes, len(mes))

}
