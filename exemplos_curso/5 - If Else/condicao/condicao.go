package main

import "fmt"

func main() {
	fmt.Println("Condicional")

	var num1 int16 = 5
	var num2 int16 = 10

	if num2/num1 == 2 {
		fmt.Println("Resultado é par")
	} else {
		fmt.Println("Resultado não é par")
	}

}
