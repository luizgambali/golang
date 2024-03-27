package main

import "fmt"

func main() {

	tarefas := make(chan int, 45)    //canal com as tarefas
	resultados := make(chan int, 45) //canal com os resultados

	//posso colocar quantas chamadas eu quiser para acelerar
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	//carrega as tarefas
	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	//fecha a fila de tarefas
	close(tarefas)

	//recebe os resultados e imprime
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	//fecha a fila de resultados
	close(resultados)
}

//especifico os canais aqui; tarefas é um canal só de entrada, resultados é um canal só de saidas
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- Fibonacci(numero)
	}
}

func Fibonacci(posicao int) int {

	if posicao <= 1 {
		return posicao
	} else {
		return Fibonacci(posicao-1) + Fibonacci(posicao-2)
	}
}
