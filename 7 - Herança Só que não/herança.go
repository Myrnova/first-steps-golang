package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    //coloca o nome do outro struct sem passar um tipo pra ele, isso seria a "herança" do Go, estudante agora terá todos os campos de pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{nome: "João", sobrenome: "Pedro", idade: 20, altura: 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"} // para ter a pessoa dentro de estudante você precisa criar um objecto da struct pessoa e passar para ele
	e2 := estudante{pessoa{nome: "João", sobrenome: "Pedro", idade: 20, altura: 178}, "Engenharia", "USP"}

	fmt.Println(e1, e2)

}
