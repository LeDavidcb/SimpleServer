package types

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Email    string
	Password string
	jwt.RegisteredClaims
}
