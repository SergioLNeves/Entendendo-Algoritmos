package PesquisaBinaria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PesquisaBinaria(t *testing.T) {

	t.Run("Should Return Number Ok", func(t *testing.T) {
		numberExpected1 := 5

		result1, _ := PesquisaBinaria(0, 50, numberExpected1)

		assert.Equal(t, numberExpected1, result1)

		numberExpected2 := 50

		result2, _ := PesquisaBinaria(0, 50, numberExpected2)

		assert.Equal(t, numberExpected2, result2)
	})

	t.Run("Should Return Error Number", func(t *testing.T) {
		numberExpected1 := 51

		result1, _ := PesquisaBinaria(0, 50, numberExpected1)

		assert.Equal(t, numberExpected1, result1)

		numberExpected2 := -1

		result2, _ := PesquisaBinaria(0, 50, numberExpected2)

		assert.Equal(t, numberExpected2, result2)
	})
}
