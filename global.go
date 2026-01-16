package simpleserver

import (
	"net/http"

	"ledavid.com/SimpleServer/types"
)

var SMux http.ServeMux = *http.NewServeMux()

var RefreshTokenTable = types.RefreshDatabase{}
