package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	secretKey := "secret-key"
	password := "I am a password"
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(password))
	result := h.Sum(nil)
	fmt.Println(hex.EncodeToString(result))
}
