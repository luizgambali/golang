package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int8
}

//passagem de parametro por referencia
func ParametroPorReferencia(valor *string) {
	*valor = "O valor foi alterado dentro da função"
}

//método tipado: ao identificar o método, a struct passa a ter o metodo
//sendo referenciado em sua estrutura
func (p Pessoa) Adicionar() {
	fmt.Println("Adicionando a pessoa no banco de dados: ")
	fmt.Println(p.nome)
	fmt.Print(p.idade)
}

//exemplo do uso de "defer"
func AtrasoNaChamada() {
	defer fmt.Println("Esta será a ultima instrução executada na funcao")

	fmt.Println("Executando o loop...")
	for i := 0; i < 10; i++ {
		fmt.Println("Elemento ", i)
	}

	//o "defer" será executado aqui, mesmo sendo declarado lá no começo da rotina

	return
}

func main() {

	nome := "Joao da silva"
	ParametroPorReferencia(&nome)
	fmt.Println(nome)

	//métodos
	p := Pessoa{"Joao da Silva", 16}
	p.Adicionar()

	AtrasoNaChamada()

	func() {
		fmt.Println("Esta é uma função anonima")
	}()

}
