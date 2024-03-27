package main

import "fmt"

func diasDaSemana(numero int) string {

	switch numero {
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
		return "Sábado"
	default:
		return "dia inválido"
	}

}

func main() {

	fmt.Println(diasDaSemana(1))
	fmt.Println(diasDaSemana(2))
	fmt.Println(diasDaSemana(3))
	fmt.Println(diasDaSemana(4))
	fmt.Println(diasDaSemana(5))
	fmt.Println(diasDaSemana(6))
	fmt.Println(diasDaSemana(7))
	fmt.Println(diasDaSemana(89))
	fmt.Println(diasDaSemana(0))
}
