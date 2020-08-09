package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinMax(t *testing.T) {
	t.Run("string min", func(t *testing.T) {
		actual := min("a", "b")
		assert.Equal(t, "a", actual)
	})

	t.Run("numeric min", func(t *testing.T) {
		actual := min("1.01", "1.02")
		assert.Equal(t, "1.01", actual)
	})

	t.Run("string max", func(t *testing.T) {
		actual := max("a", "b")
		assert.Equal(t, "b", actual)
	})

	t.Run("numeric max", func(t *testing.T) {
		actual := max("1.01", "1.02")
		assert.Equal(t, "1.02", actual)
	})
}
