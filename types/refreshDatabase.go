package types

import (
	"fmt"

	"ledavid.com/SimpleServer/misc"
)

type RefreshDatabase map[RefreshTokensRelation]bool

type RefreshTokensRelation struct {
	ClientName   string
	RefreshToken string
}

func (self *RefreshDatabase) Exists(token *string) bool {
	for k := range *self {
		if k.RefreshToken == *token {
			return true
		}
	}
	return false
}

func (self *RefreshDatabase) Add(token, client string) {
	relation := RefreshTokensRelation{ClientName: client, RefreshToken: token}
	(*self)[relation] = true
}

// This function tries to Regenerate a token that's in the RefreshDatabase, if the token is not found, it will return a error.
func (self *RefreshDatabase) Regenerate(token string) (string, error) {
	for k := range *self {
		if k.RefreshToken == token {
			rt, err := misc.GenerateRamdomToken()
			if err != nil {
				return "", fmt.Errorf("Error while generating the ramdom token")
			}
			k.RefreshToken = rt
			return rt, nil
		}
	}
	return "", fmt.Errorf("Coudln't find the token")
}

func (self *RefreshDatabase) Delete(token string) {
	for k := range *self {
		if k.RefreshToken == token {
			delete(*self, k)
			break
		}
	}
}
func (self *RefreshDatabase) Generate() (string, error) {
	var token string
	var err error
	for {
		token, err = misc.GenerateRamdomToken()
		if err != nil {
			fmt.Println("Error while generating the token", err)
			return "", err
		}
		exists := self.Exists(&token)
		if !exists {
			break
		}
	}
	return token, nil

}
