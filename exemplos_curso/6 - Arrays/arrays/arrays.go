package main

import "fmt"

func ImprimeArray(dias [7]string) {

	for i := 0; i < 7; i++ {
		fmt.Println((dias[i]))
	}

}

func main() {

	//array com inicializacao
	dia := [7]string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sabado"}

	ImprimeArray(dia)

	//array sem inicializacao
	var dias [7]string

	dias[0] = "Domingo"
	dias[1] = "Segunda"
	dias[2] = "Terça"
	dias[3] = "Quarta"
	dias[4] = "Quinta"
	dias[5] = "Sexta"
	dias[6] = "Sabado"

	ImprimeArray(dias)
}
