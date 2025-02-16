package Hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Hash(t *testing.T) {
	t.Run("Should Return Hash with Success", func(t *testing.T) {
		value := Hash("leite")
		assert.Equal(t, 4.99, value)

	})
}
