package main

import (
	"fmt"
	"math"
)

//? INTERFACES são usadas quando precisamos certa flexibilidades com os tipos

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type forma interface { // define que as structs precisar ter um método cuja a assinatura é igual a area() float64
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área d forma é %0.2f\n", f.area())

}

func main() {
	r := retangulo{10, 15}

	escreverArea(r)

	c := circulo{10}

	escreverArea(c)

}
