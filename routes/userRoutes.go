package routes

import (
	"TaxiServiceBackend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	//create users
	router.POST("/users/:user_type", controllers.CreateUser)
	//get all customers
	router.GET("/users/:user_type", controllers.GetAllCustomers)
}
