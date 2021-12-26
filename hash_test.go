package gopass

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPasswordNoError(t *testing.T) {
	hash, err := HashPassword("Password")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	fmt.Println(*hash)
}
