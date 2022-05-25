package main

import "fmt"

var n int

func init() { //? é executada antes da função main
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")

	fmt.Println(n)

}
