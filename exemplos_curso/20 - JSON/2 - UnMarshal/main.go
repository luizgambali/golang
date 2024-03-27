package main

import (
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
	cachorroEmJSon := `{"nome":"Rex","raca":"Dalmata","idade":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSon), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorroEmJSon2 := `{"nome":"Tobby","raca":"Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroEmJSon2), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
