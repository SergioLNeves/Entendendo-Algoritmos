package Soma

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Soma(t *testing.T) {
	t.Run("Should Return SomaResult with success", func(t *testing.T) {
		list := []int{
			1, 2, 3,
		}

		result := Soma(list)
		assert.Equal(t, 6, result)
	})

}
