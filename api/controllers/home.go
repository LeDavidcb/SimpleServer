package controllers

import (
	"io"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello world\n")
}
