package util

import "golang.org/x/crypto/bcrypt"

func Password(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(pwd, vPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(vPwd), []byte(pwd)); err != nil {
		return false
	}
	return true
}
