package main

import "fmt"

//? interface {} faz com que tudo atenda os requisitos dessa interface, sendo um tipo genérico
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica(12)
}
