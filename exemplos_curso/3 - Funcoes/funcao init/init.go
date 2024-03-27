package main

import "fmt"

func init() {
	fmt.Println("Texto impresso pela função init. Sempre é executara antes!")
}
func main() {
	fmt.Println("Texto impresso pela main. A init é executada antes...")
}
