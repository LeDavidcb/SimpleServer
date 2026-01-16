package middlewares

import (
	jwt1 "ledavid.com/SimpleServer/jwt"
	"net/http"
	"strings"
)

func JwtAuth(target http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		v, ok := r.Header["Authorization"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No Auth jwt was specified"))
			return
		}

		jwtToken := strings.Split(v[0], " ")[1]

		token, err := jwt1.ParseJwt(jwtToken)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Token is not longer valid, Please refresh it"))
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Token is not valid"))
			return
		}
		target(w, r)
	}
}
