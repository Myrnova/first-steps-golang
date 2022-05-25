package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int) { //estou dizendo que a função tem um retorno com nome ja definido
	soma = n1 + n2 //não precisa instanciar as variáveis já que já foram ali em cima
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
