package main

import "fmt"

//sequencia de fibonacci, o proximo número é sempre a soma dos dois números anteriores
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}

func main() {

	fmt.Println("Funções Recursivas")

	posicao := uint(12)

	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
