package PesquisaBinaria

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PesquisaBinaria(t *testing.T) {
	list := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	}

	t.Run("Should Return Number Ok", func(t *testing.T) {
		numberExpected1 := 5

		result1, err := PesquisaBinaria(list, numberExpected1)
		assert.Nil(t, err)
		assert.Equal(t, numberExpected1, result1)

		numberExpected2 := 49

		result2, err := PesquisaBinaria(list, numberExpected2)
		assert.Nil(t, err)
		assert.Equal(t, numberExpected2, result2)
	})

	t.Run("Should Return Error Number", func(t *testing.T) {
		numberExpected := -1
		err := errors.New("Item não encontrado")

		number, err := PesquisaBinaria(list, numberExpected)
		assert.NotNil(t, err)
		assert.Equal(t, numberExpected, number)
		assert.Error(t, err, "Item não encontrado")
	})
}
