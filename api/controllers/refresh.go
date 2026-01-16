package controllers

import (
	"io"
	"net/http"
	"strings"

	simpleserver "ledavid.com/SimpleServer"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	rt := r.Header["Refresh-Token"][0]
	strings.Trim(rt, " ")
	newrt, err := simpleserver.RefreshTokenTable.Regenerate(rt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.WriteString(w, newrt)
}
