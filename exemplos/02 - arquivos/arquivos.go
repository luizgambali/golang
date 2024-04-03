package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var p = fmt.Println

func main() {

	//REFERENCIA: https://www.devdungeon.com/content/working-files-go#get_file_info

	p("Criando um arquivo vazio...")
	CriarArquivoVazio()

	p("Verificando se o arquivo existe")
	ArquivoExiste("teste.txt")

	if ArquivoExiste("teste123.txt") {
		p("Excluindo arquivo teste123.txt")
		ExcluirArquivo("teste123.txt")
	}

	p("Renomeando o arquivo teste.txt para teste123.txt")
	RenomearArquivo("teste.txt", "teste123.txt")

	p("Escrevendo no arquivo")

	Escrever("teste123.txt", "Este é um texto escrito diretamente no arquivo!")
	conteudo := Ler("teste123.txt")
	p("Conteudo do arquivo: ", conteudo)
}

func CriarArquivoVazio() {
	p("Criando um arquivo vazio")
	novoArquivo, err := os.Create("teste.txt")

	if err == nil {
		p(novoArquivo)
	} else {
		p("Falha ao criar arquivo")
	}

	novoArquivo.Close()
}

func ArquivoExiste(nomeArquivo string) bool {
	_, err := os.Stat(nomeArquivo)

	if err != nil {
		if os.IsNotExist(err) {
			p("O arquivo não existe")
		}
		return false
	} else {
		return true
	}
}

func RenomearArquivo(nomeArquivo string, novoNome string) {

	if !ArquivoExiste(nomeArquivo) {
		p("O arquivo '", nomeArquivo, "' não existe")
	} else if ArquivoExiste(novoNome) {
		p("Já existe um arquivo com o novo nome: ", novoNome)
	} else {
		err := os.Rename(nomeArquivo, novoNome)

		if err != nil {
			p("Falha ao renomear o arquivo", err)
		} else {
			p("Arquivo renomeado com sucesso")
		}
	}
}

func ExcluirArquivo(nomeArquivo string) bool {
	if ArquivoExiste(nomeArquivo) {
		os.Remove(nomeArquivo)
		return true
	} else {
		return false
	}
}

func PermissaoLeitura(nomeArquivo string) bool {
	arquivo, err := os.OpenFile(nomeArquivo, os.O_RDONLY, 0666)

	if err != nil {
		if os.IsPermission(err) {
			p("Erro: não é possível ler o arquivo (denied)")
		}
		return false
	}

	arquivo.Close()
	return true
}

func PermissaoEscrita(nomeArquivo string) bool {
	arquivo, err := os.OpenFile(nomeArquivo, os.O_WRONLY, 0666)

	if err != nil {
		if os.IsPermission(err) {
			p("Erro: não é possível escrever no arquivo (denied)")
		}
		return false
	}

	arquivo.Close()
	return true
}

func Escrever(nomeArquivo string, texto string) {

	arquivo, err := os.OpenFile(nomeArquivo, os.O_WRONLY, 0666)

	defer arquivo.Close() //fecja o arquivo ao final da rotina

	if err == nil {

		buffer := bufio.NewWriter(arquivo)

		bytes, err := buffer.WriteString(texto)

		if err != nil {
			p(err)
		} else {
			p("Escrevendo ", bytes)
		}
		buffer.Flush()
		buffer.Reset(buffer)

	} else {
		if os.IsNotExist(err) {
			p("O arquivo não existe")
		} else if os.IsPermission(err) {
			p("Falha de permissão para gravar no arquivo")
		}
	}
}

func Ler(nomeArquivo string) string {

	arquivo, err := os.Open(nomeArquivo)

	if err != nil {
		p("Erro ao abrir o arquivo")
		return ""
	} else {

		defer arquivo.Close()

		conteudo, err := io.ReadAll(arquivo)

		if err == nil {
			return string(conteudo)
		} else {
			p(err)
			return ""
		}
	}
}
