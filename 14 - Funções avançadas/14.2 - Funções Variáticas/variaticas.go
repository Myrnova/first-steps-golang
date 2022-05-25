package main

import "fmt"

// função variática é uma função que pode receber n números como paramêtros

func soma(numeros ...int) int { //diz que vai receber de um a n int
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) { // não é possivel utilizar dois parametros variaticos e ele tem que obrigatoriamente ser o ultimo definido

	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11))
	escrever("Olá mundo", 10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
}
