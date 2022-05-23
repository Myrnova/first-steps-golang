package main

func main() {
	//? ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 2

	println(soma, subtracao, divisao, multiplicacao, resto)

	/*
		var numero1 int16 = 10
		var numero2 int32 = 25
		//* go não deixa você fazer operações com tipos diferentes, não é possivel somar int16 com int32
		 soma2 := numero1 + numero2
	*/

	//? FIM ARITMÉTICOS

	//? OPERADORES DE ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String 2"
	println(variavel1, variavel2)
	//? FIM DOS OPERADORES DE ATRIBUIÇÃO

	//? OPERADORES RELACIONAIS
	println(1 > 2)
	println(1 >= 2)
	println(1 == 2)
	println(1 <= 2)
	println(1 > 2)
	println(1 < 2)
	println(1 != 2)

	//? FIM OPERADORES RELACIONAIS

	//? OPERADORES LÓGICOS
	println("----------- Operadores Lógicos ----------")
	verdadeiro, falso := true, false
	println(verdadeiro && falso) // && - AND
	println(verdadeiro || falso) // || - OR
	println(!verdadeiro)         // ! - NEGATION
	//? FIM DOS OPERADORES LÓGICOS

	//? OPERADORES UNÁRIOS
	numero := 10

	numero++
	numero += 15 // numero = numero + 15

	numero--
	numero -= 15 // numero = numero - 15

	numero *= 3 // numero = numero * 3

	numero /= 10 // numero = numero / 10

	numero %= 3 // numero = numero % 3

	println(numero)
	//? FIM DOS OPERADORES UNÁRIOS

	//? OPERADOR TERNÁRIO

	var texto string

	if numero > 5 {
		texto = "Maior que cinco"
	} else {
		texto = "Menor que cinco"
	}
	// texto = numero > 5 ? "Maior que cinco" : "Menor que cinco" // não existe o ternãrio no GO
	println(texto)
	//? FIM OPERADOR TERNÁRIO
}
