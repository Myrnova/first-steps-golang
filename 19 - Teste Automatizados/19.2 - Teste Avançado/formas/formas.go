package formas

import (
	"fmt"
	"math"
)

//? INTERFACES são usadas quando precisamos certa flexibilidades com os tipos

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

type Forma interface { // define que as structs precisar ter um método cuja a assinatura é igual a area() float64
	Area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A área d forma é %0.2f\n", f.Area())

}

func main() {
	r := Retangulo{10, 15}

	EscreverArea(r)

	c := Circulo{10}

	EscreverArea(c)

}
