package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) { //*int = ponteiro de int
	*numero *= -1 //*numero seria um desponteirização
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) // & envia o endereço de memória onde a variável está salvo
	fmt.Println(novoNumero)

}

//? quando enviamos um valor sem utilizar o ponteiro, estamos enviando uma copia para a função
//? quando enviamos com paramêtros estamos passando uma referencia para essa função e qualquer alteração a referencia vai alterar a variável
