package utils

import "golang.org/x/crypto/bcrypt"

func CreateHashPassword(curPass string) string {
	password := []byte(curPass)
	hashPass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	return string(hashPass)
}

func CompareHasPassword(passUser string, curPass string) error {
	hashedPassword := []byte(passUser)
	password := []byte(curPass)
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
