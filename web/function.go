package web

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Recover() {
	if err := recover(); err != nil {
		log.Println(SERVER_RECOVER_FROM_ERROR)
		log.Println(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
