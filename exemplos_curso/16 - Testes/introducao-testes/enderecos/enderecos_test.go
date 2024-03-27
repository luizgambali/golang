package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTiposDeEnderecos(t *testing.T) {
	//nome do arquivo precisa terminar com "_test.go"
	//nome da função de testes precisa começar com Test e receber o parametro, conforme acima

	cenariosDeTeste := []cenarioDeTeste{
		{ "Rua ABC", "Rua"},
		{ "Avenida Paulista", "Avenida"},
		{ "Rodovia dos Imigrantes", "Rodovia"},
		{ "Praça das Rosas", "Tipo Inválido"},
		{ "Estrada Qualquer", "Estrada"},
		{ "RUA DAS FLORES", "Rua"},
		{ "", "Tipo Inválido" },
		{ "AVENIDA REBOUÇAS", "Avenida" }, 
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
					tipoDeEnderecoRecebido,
					cenario.retornoEsperado,
					)
		}
	}
}
