package jwt

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"ledavid.com/SimpleServer/types"
)

func GetJwtSecret() (string, error) {
	jsecret := os.Getenv("JWTSECRET")
	if jsecret == "" {
		return "", fmt.Errorf("Error while tring to get the env variable for the jwt secret")
	}
	return jsecret, nil
}

func GenerateJwt(claims *types.UserClaims) (string, error) {
	secret, err := GetJwtSecret()
	if err != nil {
		return "", fmt.Errorf("Falied to get the jwt secret: %s", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	jwt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("Falied to sign the jwt: %s", err)
	}
	return jwt, nil
}
func ParseJwt(signedJwt string) (*jwt.Token, error) {

	secret, err := GetJwtSecret()
	if err != nil {
		return &jwt.Token{}, fmt.Errorf("Falied to get the jwt secret: %s", err)
	}
	parsedJwt, err := jwt.Parse(signedJwt, func(t *jwt.Token) (any, error) { return []byte(secret), nil })
	if err != nil {
		return &jwt.Token{}, fmt.Errorf("Failed to parse jwt: %s", err)
	}
	return parsedJwt, nil
}
