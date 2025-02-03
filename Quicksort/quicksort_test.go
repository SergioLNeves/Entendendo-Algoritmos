package Quicksort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_QuickSort(t *testing.T) {
	lista := []int{10, 5, 2, 3}
	t.Run("QuickSort With Success", func(t *testing.T) {
		listExpected := []int{2, 3, 5, 10}
		list := QuickSort(lista)
		assert.Equal(t, list, listExpected)
	})
}
