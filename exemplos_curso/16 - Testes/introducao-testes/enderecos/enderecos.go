package enderecos

import "strings"

// TipoDeEndereco: Verifica se tem um tipo de endereço válido e retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	} else {
		return "Tipo Inválido"
	}
}
