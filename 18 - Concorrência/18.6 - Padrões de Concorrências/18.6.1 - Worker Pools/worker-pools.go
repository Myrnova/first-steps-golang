//? worker pools é uma fila de tarefa a ser executada e você tem vários workers que pegam essas tarefas para serem executadas de forma independente

package main

import "fmt"

//sequencia de fibonacci, o proximo número é sempre a soma dos dois números anteriores
func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}

func worker(tarefas <-chan int, resultados chan<- int) { //<-chan significa que só recebe dados, enquanto chan<- que só envia dados
	for tarefa := range tarefas {
		resultados <- fibonacci(tarefa)
	}
}

func main() {
	tarefas := make(chan int, 45)

	resultados := make(chan int, 45)

	// chega no for do worker e fica esperando vir uma tarefa, então cada um está esperando uma tarefa e quando vem, chama o fibonnaci, cada worker está recebendo um tarefa
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}
