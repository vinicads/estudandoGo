package enderecos

import "testing"

type cenarioDeTeste struct {
	endereco string
	tipoDeEnderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"AVENIDA REBOUÇAS", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {

	tipoDeEnderecoRecebido := TipoDeEndereco(cenario.endereco)

	if (cenario.tipoDeEnderecoEsperado != tipoDeEnderecoRecebido){
		t.Errorf("O tipo recebido é diferente do esperado. Esperava %s e recebeu %s", cenario.tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}
}
