//? metodos são funções dentro de estruturas

package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	//salvar func() // aqui está sendo dito que tem um campo salvar que é do tipo func()
}

func (u usuario) salvar() { // criando uma função que está "grudada" a uma certa estrutura
	fmt.Printf("Salvando os dados do Usuário %s de idade %d no banco de dados\n", u.nome, u.idade)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() { //para podermos alterar um dado dentro do struct passado, precisamos adicionar o ponteiro (*) pois assim conseguimos persistir o dado alterado
	u.idade++ //não precisa fazer a desreferenciação
}

func main() {

	usuario1 := usuario{"Usuario 1", 20}

	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}

	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()

	fmt.Println((maiorDeIdade))

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
