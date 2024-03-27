package main

import (
	"arquivos/dados"
	"arquivos/utils"
	"fmt"
)

func main() {

	LerGravarDados()
	OperacaoDiretorio()

	// copy file
	fmt.Println("Copiando o arquivo...")
	utils.CopiarArquivo("data.txt", "data2.txt")

}

func LerGravarDados() {

	fmt.Println("Lendo dados do arquivo")

	text := dados.LerDados()

	fmt.Println(text)

	fmt.Println("Escrevendo dados no arquivo")
	dados.GravarDados("Ola mundo!")

	fmt.Println("Adicionando dados ao arquivo")
	dados.AdicionarDados("Ola mundo novamente!")
}

func OperacaoDiretorio() {

	fmt.Println("Criando um novo diretório...")
	utils.CriarPasta("novapasta")

	fmt.Println("Removendo um diretório...")
	utils.RemoverPasta("novapasta")

	fmt.Println("Retornando a lista de arquivos de um diretorio './'")
	utils.MostrarArquivos("./")

	fmt.Println("Retornando somente o nome dos arquivos do diretorio './'")
	fileNames := utils.ListarArquivos("./")

	fmt.Printf("\t%v\n", fileNames)
}
