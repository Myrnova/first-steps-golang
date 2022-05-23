package main

import (
	"errors"
	"fmt"
)

// no Go todo tipo de dados tem o seu valor inicial, para string é um espaço em branco, para int ou float é o 0, para bool é false e nil para error

func main() {

	//? NÚMEROS INTEIROS
	//existe int8 int 16 int32 int64, e caso não especifique int pegara a arquitetura do computador
	var numero int = 1000000000000
	fmt.Println(numero)
	// apenas números sem sinais, segue a mesma premissa do int, uint8, ...
	var numero2 uint32 = 100
	fmt.Println(numero2)
	//? FIM NÚMEROS INTEIROS

	//? ALIAS
	// INT32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)
	//? FIM ALIAS

	//? NÚMEROS REAIS
	// não é possível criar utilizando float, mas se utilizar inferência ele sai como float que pega a arquitetura do computador
	var numeroReal1 float32 = 123.45
	var numeroReal float64 = 130000000.45
	numeroReal3 := 12345.67
	fmt.Println(numeroReal1, numeroReal, numeroReal3)
	//? FIM NÚMEROS REAIS

	//? STRINGS
	var str string = "Texto" //sempre com aspas duplas
	fmt.Println(str)
	str2 := "Texto2" //inferência
	fmt.Println(str2)
	//não existe char no Go
	char := 'B' //pega o número desse caracterer na tabela ASC
	fmt.Println(char)
	//? FIM STRINGS

	var texto int16
	fmt.Println(texto)

	//? BOOL
	var booleano1 bool = true
	fmt.Println(booleano1)
	//? FIM BOOL

	//? ERROR
	var erro error = errors.New("Erro interno") //errors é um pacote nativo para criar erros
	fmt.Println(erro)
	//no go existe um tipo especifico de tipo que é o erro
	//? FIM ERRO

}
