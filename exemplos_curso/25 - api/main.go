package main

import (
	"fmt"
	"log"
	"net/http"
	"webapi/src/router"
	"webapi/src/teste"
)

func main() {

	r := router.Gerar()
	t := teste.TestePacote() //só para testar a importação do pacote

	fmt.Println(t)

	log.Fatal(http.ListenAndServe(":5000", r))
}
