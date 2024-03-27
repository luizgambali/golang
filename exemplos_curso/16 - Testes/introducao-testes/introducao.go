package main

import (
	"Aprendizado/15-Testes/introducao-testes/enderecos"
	"fmt"
)

func main() {

	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")

	fmt.Println(tipoEndereco)
}
