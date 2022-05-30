package main

import (
	"fmt"
	"time"
)

func main() {

	go escrever("Olá mundo") //goroutine
	// go escrever("Programando em GO!") // nesse caso o programa não vai imprimir nada na tela pois vai passar pela primeira e dizer "nao vou esperar essa acabar"
	// e passa para a próxima que também é uma goroutine então também não espera acabar e chega no final do programa e acaba a execução
	escrever("Programando em GO!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

// paralelismo acontece quando duas ou mais tarefas são executadas ao mesmo tempo, acontece quando tem um processador com mais de um núcleo elas são divididas entre eles
// concorrência duas tarefas estariam rodando, mas não espera a outra acabar para rodar, elas dividem o tempo, roda um pouco da primeira para roda a segunda para e por ai vai
