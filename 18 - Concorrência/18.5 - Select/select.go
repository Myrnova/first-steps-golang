package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		//? caso o canal 1 possa receber a mensagem, printa na tela, caso o canal 2 possa receber na tela, printa na tela também
		//? não faz um ficar esperando o outro
		select {
		case mensagem1 := <-canal1:
			fmt.Println(mensagem1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}

		// mensagemCanal1 := <-canal1 //espera 0.5 segundos para receber a mensagem
		// // fica esperando o segundo canal receber mensagem para começar a receber novamente no canal 1, sendo que podia imprimir 4 vezes

		// mensagemCanal2 := <-canal2 //espera 2 segundos para receber a mensagem
		// fmt.Println(mensagemCanal1)
		// fmt.Println(mensagemCanal2)
	}
}
