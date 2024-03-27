package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) MostrarDados() {
	fmt.Println(c.nome)
}

func main() {

	mensagem("Ola Mundo!")

	fmt.Println(add(42, 13))

	soma, subtracao := operations(42, 13)

	fmt.Printf("Soma: %d, Subtracao: %d\n", soma, subtracao)

	totalEmprestimo, parcelas := CalcularJurosSimples(1000, 0.1, 5)

	fmt.Printf("Total do emprestimo: %.2f, Parcelas: %.2f\n", totalEmprestimo, parcelas)

	cliente := Cliente{nome: "Fulano"}

	cliente.MostrarDados()

	f := func() {

		fmt.Println("Funcao anonima")

	}

	f()

}

func mensagem(msg string) {
	fmt.Println(msg)
}

func add(x int, y int) int {
	return x + y
}

func operations(x int, y int) (int, int) {
	return x + y, x - y
}

func CalcularJurosSimples(valorEmprestimo float32, taxaJuros float32, prazo uint) (float32, float32) {

	totalEmprestimo := (valorEmprestimo * taxaJuros) * float32(prazo)

	parcelas := totalEmprestimo / float32(prazo)

	return totalEmprestimo, parcelas
}
