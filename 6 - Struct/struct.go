package main

import "fmt"

//Struct = Classe
//quando você inicializa um struct e não passa o valor ele coloca todos os campos com seu respectivo valor zero
//https://www.golangprograms.com/go-language/struct.html
type usuario struct {
	nome     string `default:"Teste"`
	idade    uint8  `default:"18"`
	endereco endereco
}

// adicionar metodo dentro do struct, consigo acessar as propriedades utilizando o "u"
func (u *usuario) maiorDeIdade() bool {
	if u.idade > 18 {
		return true
	}
	return false
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21

	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2.maiorDeIdade())

	usuario3 := usuario{nome: "Davi"}

	fmt.Println(usuario3)

}
