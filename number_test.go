package gopass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsEvenTrue(t *testing.T) {
	assert.True(t, isEven(4))
}

func Test_IsEvenFalse(t *testing.T) {
	assert.False(t, isEven(5))
}
