package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		rest := Retangulo{10, 15}
		areaEsperada := float64(150)
		areaRecebida := rest.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("Area recebida %.2f é diferenet da area esperada %.2f", areaRecebida, areaEsperada)
			//Errorf se encontrar um erro ele mostra o erro e continua, fatalf para os testes nele
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()
		if areaEsperada != areaRecebida {
			t.Fatalf("Area recebida %.2f é diferenet da area esperada %.2f", areaRecebida, areaEsperada)
			//Errorf se encontrar um erro ele mostra o erro e continua, fatalf para os testes nele
		}
	})
}
