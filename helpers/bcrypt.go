package helpers

import "golang.org/x/crypto/bcrypt"

func HasPass(p string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), 8)
	return string(hash)
}

func ComparePass(hashPassword, passwordUSer []byte) bool {
	hash, password := []byte(hashPassword), []byte(passwordUSer)
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
