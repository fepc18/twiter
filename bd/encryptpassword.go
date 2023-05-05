package bd

import "golang.org/x/crypto/bcrypt"

// EncryptPassword is the function that allows to encrypt the password
func EncryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
