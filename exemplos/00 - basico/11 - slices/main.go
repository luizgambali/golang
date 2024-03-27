package main

import (
	"fmt"
	"slices"
)

var nomes []string

func main() {

	AdicionarItens()
	MostrarItens()

	RemoverItens("Maria")
	ListarItens()
	AtualizarItens("José", "José da Silva")
	ListarItens()

	//a partir do go 1.8
	fmt.Println("Contém o item 'João'? ", slices.Contains(nomes, "João"))

	fmt.Println("Removendo dois itens, começando no indice 0")
	var result = slices.Delete(nomes, 0, 2)
	fmt.Println(result)

}

func AdicionarItens() {

	nomes = append(nomes, "João")
	nomes = append(nomes, "Maria")
	nomes = append(nomes, "José")
	nomes = append(nomes, "Pedro")
}

func RemoverItens(nome string) {

	fmt.Println("Removendo o nome: ", nome)

	for i, v := range nomes {
		if v == nome {
			nomes = append(nomes[:i], nomes[i+1:]...)
		}
	}
}

func AtualizarItens(nome string, novoNome string) {

	fmt.Println("Atualizando o nome: ", nome)
	for i, v := range nomes {
		if v == nome {
			nomes[i] = novoNome
		}
	}
}

func ListarItens() {
	fmt.Println(nomes)
}

func MostrarItens() {

	fmt.Println("nomes ", nomes)
	fmt.Println("nomes[0] ", nomes[0])
	fmt.Println("nomes[:1] ", nomes[:1])
	fmt.Println("nomes[1:] ", nomes[1:])
	fmt.Println("nomes[1:2] ", nomes[1:2])
	fmt.Println("nomes[:] ", nomes[:])

}
