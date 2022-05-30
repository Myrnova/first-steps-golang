package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //quantas 	goroutines vai ter rodando ao mesmo tempo

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() //retira um cara do contador adicionado a cima // -1
	}()
	go func() {
		escrever("Programando em GO!")
		waitGroup.Done() // -1
	}()
	go func() {
		escrever("Goroutine 3!")
		waitGroup.Done() // -1
	}()
	go func() {
		escrever("Goroutine 4!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() //fala pra função main esperar a contagem da goroutine zerar
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
