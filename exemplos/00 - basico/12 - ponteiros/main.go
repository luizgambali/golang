package main

func main() {

	println("1. Criando a variavel 'a' do tipo inteiro, e atribuindo o valor 10.")
	var a int = 10

	println("\n2. Vamos imprimir o endereço de 'a', e o valor de 'a'.")
	println("Endereço de a:", &a)
	println("Valor de a:", a)

	println("\n3. Vamos criar um ponteiro b, que aponta para o endereço de a")
	var b *int = &a

	println("Endereço de b:", &b)
	println("Valor de b (lembrando que b é um ponteiro de inteiro, apontando para o endereço de a):", b)
	println("Valor de *b:", *b)
	println("B: ", b)

	println("\n4. Vamos criar um ponteiro c, que aponta para o endereço de a")
	var c *int = &a

	println("\n5. Vamos alterar o valor de a, e conferir se c também foi alterado")
	a = 15

	println("Valor de a:", a)
	println("Valor de *b:", *b)
	println("Valor de *c:", *c)

	println("\n6. Agora, vamos alterar o valor de c, e ver se a também foi alterado")

	*c = 20
	println("Valor de a:", a)
	println("Valor de *b:", *b)
	println("Valor de *c:", *c)

	println("\n7. Vamos chamar uma função e passar o valor de a, e ver se a b e c foram alterados")
	println("Endereco de a, antes de chamar a função: ", &a)
	AlterarValor(a)

	println("Valor de a:", a)
	println("Valor de *b:", *b)
	println("Valor de *c:", *c)

	println("\n8. Vamos chamar uma função passando a, e dentro da função tentar acessar o endereço e alterar o valor")
	println("Endereco de a, antes de chamar a função: ", &a)
	AlterarValor1(a)

	println("Valor de a:", a)
	println("Valor de *b:", *b)
	println("Valor de *c:", *c)

	println("\n9. Vamos chamar uma função passando o endereço de a, e dentro da função tentar alterar o valor de a")
	println("Endereco de a, antes de chamar a função: ", &a)
	AlterarValor2(&a)

	println("Valor de a:", a)
	println("Valor de *b:", *b)
	println("Valor de *c:", *c)

}

func AlterarValor(a int) {
	a = 12
	println("")
	println("\tEndereço de a, dentro da funcao:", &a)
	println("\tValor de a, dentro da funcao:", a)
	println("")
}

func AlterarValor1(a int) {
	var x *int = &a

	*x = 35
	println("")
	println("\tValor de a, dentro da função:", a)
	println("\tEndereço de a, dentro da função:", &a)
	println("\tEndereco de x, dentro da função:", &x)
	println("\tValor de *x, dentro da função:", *x)
	println("")
}

func AlterarValor2(a *int) {
	x := a

	*x = 55
	println("")
	println("\tValor de a, dentro da funcao:", a)
	println("\tValor de x, dentro da função:", *x)
	println("\tEndereco de x, dentro da função:", &x)
	println("\tValor de *x:", *x)
	println("")
}
