package main

import "fmt"

func Fibonacci(posicao uint) uint {

	if posicao <= 1 {
		return posicao
	} else {
		return Fibonacci(posicao-1) + Fibonacci(posicao-2)
	}
}

func main() {

	fmt.Println(Fibonacci(10))

}
