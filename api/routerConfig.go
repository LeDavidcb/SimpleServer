package api

import "ledavid.com/SimpleServer/misc"
import "ledavid.com/SimpleServer/api/controllers"

func AddRoutesAndMiddlewares() {
	// All middlewares
	// All routes
	misc.SMux.HandleFunc("/", controllers.Home)

}
