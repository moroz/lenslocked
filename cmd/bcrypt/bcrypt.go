package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func comparePassword(hash []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func main() {
	password := "example"
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashed))

	passwords := []string{"example", "invalid"}

	for _, pass := range passwords {
		if comparePassword(hashed, pass) {
			fmt.Printf("Password is valid: %s\n", pass)
		} else {
			fmt.Printf("Password is invalid: %s\n", pass)
		}
	}
}
