package main

import (
	"errors"
	"fmt"
)

// funcao sem parametro, e sem retorno
func FuncaoSemRetornoSemParametro() {
	fmt.Println("Esta função não possui parametros nem retorno")
}

// funcao com parametro, mas sem retorno
func FuncaoComParametro(texto string) {
	fmt.Println(texto)
}

// funcao com parametro, e com retorno
func Soma(num1, num2 int) int {
	return num1 + num2
}

// funcao com mais de um retorno
func TesteRetorno(num1, num2 int) (int, string) {
	return num1 + num2, "este valor é um retorno!"
}

func RetornoNomeado() (chave string, valor string) {

	//se nomear o retorno, não é necessário declarar a variavel

	chave = "Retorno para o parametro chave"
	valor = "Retorno para o parametro valor"

	return chave, valor
}

//funcao que retorna um erro

func RetornoComErro() (valor string, err error) {
	return "Teste de retorno com erro", errors.New("mensagem de erro retornada da função")
}

func main() {

	FuncaoSemRetornoSemParametro()
	FuncaoComParametro("Esta funcao recebeu como parametro este texto!")
	fmt.Println("10 + 20: ", Soma(10, 20))

	soma, retorno := TesteRetorno(10, 10)

	fmt.Println("Soma dos numeros: ", soma)
	fmt.Println("Retorno: ", retorno)

	//ignora o segundo parametro, pega apenas o primeiro valor
	soma, _ = TesteRetorno(10, 10)
	fmt.Println("Soma dos numeros, quando ignorado o segundo parametro: ", soma)

	mensagem, erro := RetornoComErro()

	if erro != nil {
		fmt.Println("Ocorreu um eror: ", erro)
	} else {
		fmt.Println("Não ocorreu erro nenhum!")
	}

	fmt.Println(mensagem)

	chave, valor := RetornoNomeado()

	fmt.Println("Chave: ", chave)
	fmt.Println("Valor: ", valor)

}
