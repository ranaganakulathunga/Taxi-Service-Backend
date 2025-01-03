package main

import (
	"TaxiServiceBackend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	//initialize gin port
	configPort := "8000"

	//create gin instance
	router := gin.New()

	//User Routes
	routes.UserRoutes((router))

	//Run gin Server
	router.Run(":" + configPort)

}
