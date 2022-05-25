package main

import "fmt"

//? Recuperar o programa pro panic não matar ele
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()

	fmt.Println("Entrando na função de aprovado")

	defer fmt.Println("Media Calculada o resultado será retornado")
	if media := (n1 + n2) / 2; media > 6 {
		//fmt.Println("Media Calculada o resultado será retornado")
		return true
	} else if media < 6 {
		return false
	}
	//? O programa entra em pânico e mata a execução do programa
	panic("A MEDIA É EXATAMENTE 6!")
	//fmt.Println("Media Calculada o resultado será retornado")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução")
}
