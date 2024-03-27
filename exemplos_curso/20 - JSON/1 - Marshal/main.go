package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` //faz o mapeamento do json
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

//caso não queira mapear um dos campos do struct, basta colocar um "-" no mapeamento,
//o go não vai mapea-lo na conversão

func main() {
	c := cachorro{"Rex", "Dalmata", 3}

	fmt.Println(c)

	cachorroEmJson, err := json.Marshal(c)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(cachorroEmJson)                  //traz um slice de bytes
		fmt.Println(bytes.NewBuffer(cachorroEmJson)) //converte o slice de bytes em algo legível
	}

	c2 := map[string]string{
		"nome": "toby",
		"raca": "pooble",
	}

	cachorroEmJson2, err := json.Marshal(c2)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(cachorroEmJson2)                  //traz um slice de bytes
		fmt.Println(bytes.NewBuffer(cachorroEmJson2)) //converte o slice de bytes em algo legível
	}
}
