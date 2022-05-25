package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 10 { // funcionando como um while
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j += 5 { // funcionando como um for
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := []string{"João", "Davi", "Lucas"}
	//range itera os elementos em uma variedade de tipos de estruturas
	for indice, nome := range nomes { // funcionando como map javascript
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" { //retorna os numeros da letra na tabela ASCII
		fmt.Println(indice, string(letra))
	}
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}
