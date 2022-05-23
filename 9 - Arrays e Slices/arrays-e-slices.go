package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Arrays e Slices")

	var array1 [5]string //? [] com quantos itens dentro do array e o tipo, não é possível fazer arrays com tipos diferentes
	array1[0] = "Posição 1"
	fmt.Println(array1)
	array2 := [5]string{"Posição 2", "Posição 3", "Posição 4", "Posição 5", "Posição 6"}
	fmt.Println(array2)
	array3 := [...]int{1, 2, 3, 4, 5} //? ... fixa o tamanho do array dependendo de quantos itens coloquei na declaração do array
	fmt.Println(array3)

	slice := []int{1, 2, 3, 4, 5, 6, 7} //? no slice não preciso passar o tamanho que ele terá, ele terá um tamanho dinâmico

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(array3))

	slice = append(slice, 18)

	fmt.Println(slice)

	slice2 := array2[1:3] //? pega a partir do 1 incluindo a posição e para no 3 sem incluir no resultado
	fmt.Println(slice2)
	array2[1] = "Posição Alterada"

	fmt.Println(slice2) // funciona como um ponteiro que aponta para um array

	// Arrays Internos
	//.make() aloca espaço na memória
	fmt.Println("------------Arrays Internos------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)
	//? o GO quando percebe que você vai estourar a capacidade do seu slice, ele criar outro array com o tamanho dobrado

	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade

	fmt.Println(slice3)

	slice4 := make([]float32, 5)
	fmt.Println(slice4)

}
