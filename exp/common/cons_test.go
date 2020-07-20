package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsString(t *testing.T) {
	list := List("1", "2", "3", "4", "5")
	assert.Equal(t, "(1 2 3 4 5)", list.String())
	list = List(1, 40, 20, 999)
	assert.Equal(t, "(1 40 20 999)", list.String())
	list = List(1, 40, List(24, 33), 100, List("c", "a", "b"), 88)
	assert.Equal(t, "(1 40 (24 33) 100 (c a b) 88)", list.String())
}
