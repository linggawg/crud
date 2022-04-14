package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
