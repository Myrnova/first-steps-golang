package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string, 2) // especificando que o canal tem uma capacidade de 2
	//? canal com buffer só bloqueia quando atingir a capacidade maxima dele, não fica esperando alguém receber o valor!
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // mandando um valor pra dentro do canal
		time.Sleep(time.Second)
	}

	close(canal) // fechando o canal para parar de enviar ou receber dados
}
