package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função  1")

}

func funcao2() {
	fmt.Println("Executando a função 2")
}

//? DEFER = ADIAR
//? É executado antes dor return da função, na função abaixo posso utiliza-lo para impedir que seja necessário repetir código
func alunoEstaAprovado(n1, n2 float32) bool {

	fmt.Println("Entrando na função de aprovado")

	defer fmt.Println("Media Calculada o resultado será retornado")
	if media := (n1 + n2) / 2; media >= 6 {
		//fmt.Println("Media Calculada o resultado será retornado")
		return true
	}
	//fmt.Println("Media Calculada o resultado será retornado")
	return false

}

func main() {
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
