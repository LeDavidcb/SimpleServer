package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	simpleserver "ledavid.com/SimpleServer"
	jwt1 "ledavid.com/SimpleServer/jwt"
	"ledavid.com/SimpleServer/types"
)

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var loginRequest types.LoginRequest
	loginRequest.ParseFromRequest(&w, r)

	//GenerateJwt
	var claims types.UserClaims = types.UserClaims{}

	claims.Email = loginRequest.Email
	claims.Password = loginRequest.Password
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 5))

	jwt, err := jwt1.GenerateJwt(&claims)
	if err != nil {
		log.Println("Error while generating jwt: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "internal server error")
		return
	}
	refreshToken, err := simpleserver.RefreshTokenTable.Generate()

	if err != nil {
		log.Println("Error while generating refreshToken: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "internal server error")
		return
	}

	rawResp := types.JwtResponse{
		Jwt:          jwt,
		RefreshToken: string(refreshToken),
	}

	resp, err := json.Marshal(rawResp)

	if err != nil {
		log.Println("Error while building the json", err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "internal server error")
		return
	}
	// store the refreshToken
	client := r.Header["User-Agent"][0]
	simpleserver.RefreshTokenTable.Add(refreshToken, client)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
