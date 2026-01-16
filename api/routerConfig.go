package api

import (
	simpleserver "ledavid.com/SimpleServer"
	"ledavid.com/SimpleServer/api/controllers"
	"ledavid.com/SimpleServer/api/middlewares"
)

// This function setup all the routes for every controller and middleware
func AddRoutesAndMiddlewares() {
	// All middlewares
	simpleserver.SMux.HandleFunc("/home", middlewares.JwtAuth(controllers.Home))
	// All routes
	simpleserver.SMux.HandleFunc("/login", controllers.Login)
	simpleserver.SMux.HandleFunc("/refresh", controllers.Refresh)

}
