package main

import "fmt"

//? função closure seria como o escopo de função do javascript
func closure() func() { //vai retornar uma função
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"

	fmt.Println(texto)

	funcaoNova := closure()

	funcaoNova()
}
