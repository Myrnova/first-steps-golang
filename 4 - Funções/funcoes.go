package main

import "fmt"

func somar(n1 int8, n2 int8) int8 { // o ultimo int8 é o tipo de retorno da função
	return n1 + n2
}

// As funções podem ter mais de um retorno
// no go é possível declarar os paramêtros caso eles sejam do mesmo tipo que é especificando o tipo do ultimo
// para tipar o tipo de retorno é só colocar () e os tipos depois de especificar os paramêtros

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)

	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto da função 1")

	resultadoSoma, _ := calculosMatematicos(10, 15)
	println(resultadoSoma)
}
