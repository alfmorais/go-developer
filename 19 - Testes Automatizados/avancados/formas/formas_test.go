package formas

import "testing"

func TestArea(t *testing.T) {
	t.Parallel()
	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{10, 15}
		areaEsperada := float64(150)
		areaRecebida := retangulo.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf(
				"A area recebida (%f) nao eh igual a area esperada (%f)",
				areaRecebida,
				areaEsperada,
			)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{10}
		areaEsperada := float64(314.1592653589793)
		areaRecebida := circulo.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf(
				"A area recebida (%f) nao eh igual a area esperada (%f)",
				areaRecebida,
				areaEsperada,
			)
		}
	})
}
