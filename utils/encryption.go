package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Decrypt(hash, pwd []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, pwd)

	if err != nil {
		return false, nil
	}
	return true, nil
}
