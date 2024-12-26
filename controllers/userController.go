package controllers

import (
	"TaxiServiceBackend/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	userType := c.Param("user_type")

	//Validate user Type
	userTypes := []string{"customer", "service_provider"}
	var userTypeValid bool
	for _, s := range userTypes {
		if s == userType {
			userTypeValid = true
			break
		}
		userTypeValid = false
	}
	if !userTypeValid {
		c.JSON(helper.GetHTTPError("Invalid request parameters.", http.StatusBadRequest, c.FullPath()))
		return

	}

}
