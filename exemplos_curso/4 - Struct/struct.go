package main

import "fmt"

//declaração simples de uma struct
type Pessoa struct {
	nome  string
	idade int8
}

//declaração de uma struct que 'herda' os dados de outra. Neste caso, não precisa definir o tipo
type Estudante struct {
	Pessoa
	curso     string
	faculdade string
}

func main() {

	//criando um objeto e já inicializando
	pessoa1 := Pessoa{"Jose da silva", 28}
	estudante := Estudante{pessoa1, "Administração", "USP"}

	fmt.Println("Dados da pessoa: ", pessoa1)
	fmt.Println("Somente o nome dela: ", pessoa1.nome)
	fmt.Println("Dados do estudante: ", estudante)
	fmt.Println("Nome do estudante: ", estudante.nome)
	fmt.Println("Curso: ", estudante.curso)

	var pessoa2 Pessoa
	var estudante2 Estudante

	pessoa2.nome = "Claudio"
	pessoa2.idade = 45

	fmt.Println("Dados da outra pessoa: ", pessoa2)

	estudante2.nome = "Mauro"
	estudante2.idade = 23
	estudante2.curso = "Letras"
	estudante2.faculdade = "PUC/RJ"

	fmt.Println(estudante2)

	estudante3 := Estudante{Pessoa{"Nicolas", 19}, "Engenharia Civil", "ITA"}

	fmt.Println(estudante3)

	fmt.Println(estudante3.Pessoa)

	estudante3.Pessoa = Pessoa{"Amanda", 20}

	fmt.Println(estudante3)
}
