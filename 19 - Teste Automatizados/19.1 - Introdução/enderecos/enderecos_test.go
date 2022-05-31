// precisa ter a nomenclatura _test.go
// TESTE UNITÁRIO

package enderecos_test

import (
	. "introducao-testes/enderecos" //. diz pro go que esse é o pacote principal
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	//TestNomeDaFunçãoQueSeráTestada

	t.Parallel() // rodar os testes em paralelo
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		// {"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA Rebouças", "Avenida"},
		// {"", "Tipo Inválido"},
	}
	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				cenario.retornoEsperado,
				retornoRecebido)
		}
	}

}

// go test ./... roda todos os testes nos packages

// go test -v //roda o teste verbosamente mostrando o passo a passo
// go test --cover //mostra a cobertura de teste
// go test --coverprofile cobertura.txt // gera um arquivo co todas as informações do teste
//go tool cover --func=cobertura.txt //le o arquivo gerado pelo coverprofile e joga no terminal
//go tool cover --html=cobertura.txt //le o arquivo gerado pelo coverprofile e joga em um arquivo html mostrando as linhas que não estão cobertas

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Deu erro")
	}
}
