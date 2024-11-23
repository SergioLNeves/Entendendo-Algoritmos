package Ordenacao

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BuscaMenorIndice(t *testing.T) {
	list := []int{
		50, 20, 30, 40, 10, 5, 70, 80, 90, 100,
	}

	t.Run("Should Return Indice", func(t *testing.T) {

		indice := BuscaMenorIndice(list)
		assert.Equal(t, 5, indice)
	})
}

func Test_OrdenacaoPorSelecao(t *testing.T) {
	list := []int{
		100, 67, 124, 8, 199, 2, 589, 123, 2, 0,
	}
	t.Run("Should Return Order List", func(t *testing.T) {

		listExpected := []int{0, 2, 2, 8, 67, 100, 123, 124, 199, 589}

		listaOrdenada := OrdenacaoPorSelecao(list)
		assert.Equal(t, listExpected, listaOrdenada)

	})
}
