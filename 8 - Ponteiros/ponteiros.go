package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var variavel1 int = 10
	var variavel2 int = variavel1 // faz uma copia do valor para a variavel2

	fmt.Println(variavel1, variavel2)

	variavel2++

	fmt.Println(variavel1, variavel2)

	//? PONTEIRO É UMA REFERENCIA DE MEMÓRIA
	var variavel3 int = 100
	var ponteiro *int = &variavel3 // *int = ponteiro de int, guarda um endereço de memória de um int, para passar para o ponteiro onde a variável está na memória é só passar "&" na frente "

	fmt.Println(*ponteiro, variavel3) //*ponteiro = desreferenciação, vai no endereço de memória e pega o valor que está la dentro

	variavel3 = 150
	fmt.Println(*ponteiro, variavel3)
}
