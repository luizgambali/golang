package main

import "fmt"

type Contato struct {
	Nome     string
	Email    string
	Telefone string
}

var contatos = []Contato{}

func main() {

	opcao := 0

	for opcao != 4 {

		ExibirMenu()

		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			AdicionarContato()
		case 2:
			ListarContatos()
		case 3:
			RemoverContato()
		case 4:
			fmt.Println("Saindo...")
		}
	}
}

func ExibirMenu() {
	fmt.Println("1 - Adicionar contato")
	fmt.Println("2 - Listar contatos")
	fmt.Println("3 - Remover contato")
	fmt.Println("4 - Sair")
}

func AdicionarContato() {

	var contato Contato

	fmt.Println("Nome:")
	fmt.Scanln(&contato.Nome)

	fmt.Println("Email: ")
	fmt.Scanln(&contato.Email)

	fmt.Println("Telefone:")
	fmt.Scanln(&contato.Telefone)

	contatos = append(contatos, contato)

	fmt.Println("Contato adicionado com sucesso")
}

func ListarContatos() {
	fmt.Println(contatos)
}

func RemoverContato() {
	var nome string
	fmt.Println("Nome do contato para remoção:")
	fmt.Scanln(&nome)

	for i, contato := range contatos {
		if contato.Nome == nome {
			resultado := "ok"

			fmt.Println(resultado)
			contatos = append(contatos[:i], contatos[i+1:]...)
			fmt.Println("Contato removido com sucesso")
			return
		}

	}

	fmt.Println("Contato não encontrado")
}
