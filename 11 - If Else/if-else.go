package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")
	numero := -15

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { //declarando uma variável durante a definição do if (if init), funciona apenas no escopo do if
		fmt.Println("Número é maior que zero")
	} else if outroNumero < -10 {
		fmt.Println("Número é menor que -10")
	}
}
