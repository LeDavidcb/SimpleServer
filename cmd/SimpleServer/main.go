package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"ledavid.com/SimpleServer/api"
	"ledavid.com/SimpleServer/misc"
)

func main() {

	loadErr := godotenv.Load()
	if loadErr != nil {
		fmt.Println("Error while tring to load .env", loadErr)
		return
	}
	_, err := misc.GetJwtSecret()
	if err != nil {
		fmt.Println(err)
	}

	// Setup all routes
	api.AddRoutesAndMiddlewares()

	// TODO: change to PORT
	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", &misc.SMux)
}
