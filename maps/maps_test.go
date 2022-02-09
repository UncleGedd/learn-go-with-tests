package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "ima test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "ima test"
		assert.Equal(t, want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("quiz")
		assert.Error(t, err)
	})

	t.Run("add a word", func(t *testing.T) {
		err := dict.Add("quiz", "ima quiz")
		got, _ := dict.Search("quiz")
		assert.Equal(t, "ima quiz", got)
		assert.NoError(t, err)
	})

	t.Run("cannot overwrite word", func(t *testing.T) {
		err := dict.Add("test", "the new definition")
		assert.Error(t, err)
	})

	t.Run("delete word", func(t *testing.T) {
		dict.Add("assessment", "ima test")
		dict.Delete("assessment")
		_, err := dict.Search("assessment")
		assert.Error(t, err)
	})
}
