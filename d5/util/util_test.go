package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PolaritiesReact(t *testing.T) {
	t.Run("returns true for aA", func(t *testing.T) {
		assert.True(t, PolaritiesReact("a", "A"))
		assert.True(t, PolaritiesReact("F", "f"))
	})

	t.Run("returns false for aa", func(t *testing.T) {
		assert.False(t, PolaritiesReact("a", "a"))
	})

	t.Run("returns false for aB", func(t *testing.T) {
		assert.False(t, PolaritiesReact("a", "B"))
		assert.False(t, PolaritiesReact("A", "B"))
		assert.False(t, PolaritiesReact("a", "b"))
	})
}
