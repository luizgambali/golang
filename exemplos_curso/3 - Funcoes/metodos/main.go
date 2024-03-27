package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario '%s' no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Fulano", 20}

	fmt.Println("Usuário 1:", usuario1)

	usuario1.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()

	fmt.Println("É maior de idade?", maiorDeIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)

}
