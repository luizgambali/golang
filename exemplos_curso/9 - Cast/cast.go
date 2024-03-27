package main

import "fmt"

func ConversaoSimples() float64 {
	var num1 float32 = 10.75

	return float64(num1)
}

func ConversaoSimples2() float32 {
	var num1 float64 = 10.9999

	return float32(num1)
}

func ConversaoSimples3() int {
	var num1 float64 = 10.22222

	return int(num1)
}

func ConversaoSimples4() float32 {
	var num1 int = 10

	return float32(num1)
}

func ConversaoSimples5() string {
	var num1 int = 10
	var num2 float32 = 33.3345667

	return fmt.Sprintf("%d; %.2f", num1, num2)
}

func main() {
	fmt.Println(ConversaoSimples(), ConversaoSimples2(), ConversaoSimples3(), ConversaoSimples4(), ConversaoSimples5())
}
