package Recursividade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Factorial(t *testing.T) {
	t.Run("Should Return number factorial with success", func(t *testing.T) {
		numberExpected := 2

		result := Factorial(numberExpected)
		assert.Equal(t, numberExpected, result)
	})

}
