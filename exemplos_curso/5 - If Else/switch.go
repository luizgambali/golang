package main

import "fmt"

func DiaDaSemana(dia int8) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Dia inválido"
	}
}

func main() {
	fmt.Println(DiaDaSemana(3))
	fmt.Println(DiaDaSemana(10))

}
