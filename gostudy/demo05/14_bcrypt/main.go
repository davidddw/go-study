package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "MyDarkSecret1"

	hashedPassword := EncodePassword([]byte(password))
	fmt.Println(hashedPassword)

	// Comparing the password with the hash
	//err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(comparePassword(string(hashedPassword), password)) // nil means it is a match
}

func comparePassword(pwdHash, comPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(comPwd))
	fmt.Println(len(pwdHash))
	return err == nil
}

// EncodePassword encode the password.
func EncodePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash[:])
}
