package gopass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparePasswordFailed(t *testing.T) {
	hash, _ := HashPassword("Password")
	plain := "Password1"
	isMatch := ComparePassword(plain, *hash)
	assert.False(t, isMatch)
}

func TestComparePasswordSuccess(t *testing.T) {
	hash, _ := HashPassword("Password")
	plain := "Password"
	isMatch := ComparePassword(plain, *hash)
	assert.True(t, isMatch)
}
