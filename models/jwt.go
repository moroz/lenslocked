package models

import (
	"log"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/moroz/lenslocked/utils"
)

const ACCESS_TOKEN_VALIDITY = time.Second * 3600 * 24

var TOKEN_SIGNER = []byte(utils.RequireEnvVar("ACCESS_TOKEN_SIGNER"))

func IssueTokenForUser(user *User) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ACCESS_TOKEN_VALIDITY)),
		Subject:   strconv.Itoa(user.ID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(TOKEN_SIGNER)
}

func DecodeAccessTokenClaims(tokenString string) int {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return TOKEN_SIGNER, nil
	})

	if err != nil {
		log.Fatal(err)
		return 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := strconv.Atoi(claims["sub"].(string))
		if err != nil {
			return 0
		}
		return id
	}
	return 0
}
