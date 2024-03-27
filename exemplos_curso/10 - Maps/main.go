package main

import "fmt"

func main() {

	//dentro das chaves é o tipo da chave, fora é o tipo do valor
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"segundo":  "Pedro",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["primeiro"])

	chaves := map[string]string{
		"chave1": "valor-da-chave-1",
		"chave2": "valor-da-chave-2",
		"chave3": "valor-da-chave-3",
	}

	fmt.Println(chaves)
	fmt.Println(chaves["chave1"])

}
