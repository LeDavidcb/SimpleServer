package misc

import (
	"errors"
	"os"
)

func GetJwtSecret() (string, error) {
	jsecret := os.Getenv("JWTSECRET")
	if jsecret == "" {
		return "", errors.New("Error while tring to get the env variable for the jwt secret")
	}
	return jsecret, nil
}
