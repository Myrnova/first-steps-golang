package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

//? quando for chamar ela na função main não precisa chamar com a clausula go
//? você encapsula a chamada de uma goroutine e retorna um canal para comunicação
func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Second)
		}
	}()

	return canal
}
