package gopass

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (*string, error) {
	//Hash password
	HashByte, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	HashString := string(HashByte)
	return &HashString, err
}
