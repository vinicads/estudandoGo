package formas

import "testing"

func TestArea(t *testing.T){
	t.Run("Retangulo", func(t *testing.T){
		retangulo := Retangulo{10, 15}
		areaEsperada := float64(150)
		areaRecebida := retangulo.area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T){
		circulo := Circulo{10}
		areaEsperada := float64(314)
		areaRecebida := circulo.area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}
	})
}