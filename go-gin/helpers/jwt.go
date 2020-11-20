package helpers

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaim struct {
	UserName string `json:userName`
	Password string `json:password`
	jwt.StandardClaims
}

func GenerateToken(userName string, password string) string {
	GJWTKey := []byte("asfasdfasd")

	// GJWTKey := []byte("Dev12334")
	Claim := CustomClaim{
		UserName: userName,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 5000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim)
	ss, err := token.SignedString(GJWTKey)
	if err != nil {
		panic(err.Error())
	}
	return ss
}
