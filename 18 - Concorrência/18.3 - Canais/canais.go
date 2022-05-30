package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string) // chan de canal e o tipo, quer dizer que ele só vai poder escrever/ler strings

	go escrever("Olá mundo", canal)

	go escrever("Programando em GO!", canal)

	go escrever("Goroutine 3!", canal)

	go escrever("Goroutine 4!", canal)

	// for {
	// 	mensagem, aberto := <-canal //esperando que canal receba um valor
	// 	// a segunda variável serve para identificar se o canal está aberto ou fechado
	// 	if !aberto {
	// 		break
	// 	}
	// 	//? não espera rodar todos os for, a partir do momento que recebe um valor ele continua a execução
	// 	fmt.Println(mensagem)
	// 	//? deadlock quer dizer que seu programa está esperando algo que não vai acontecer
	// }

	for mensagem := range canal { // maneira mais simples de pegar as mensagens do canal
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // mandando um valor pra dentro do canal
		time.Sleep(time.Second)
	}

	close(canal) // fechando o canal para parar de enviar ou receber dados
}
