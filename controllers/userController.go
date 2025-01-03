package controllers

import (
	database "TaxiServiceBackend/connection"
	"TaxiServiceBackend/helper"
	"TaxiServiceBackend/models"
	"encoding/json"
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

	// Decode JSON Request Body.
	var userObject models.User
	decErr := json.NewDecoder(c.Request.Body).Decode(&userObject)
	if decErr != nil {
		c.JSON(helper.GetHTTPError("Invalid request body.", http.StatusBadRequest, c.FullPath()))
		return
	}

	// Insert data into the `test` table

	query := `
INSERT INTO customer 
(id, phone_number, user_level, gender, user_sub_type, payment_methods, wallet_amount, favorite_locations, first_name, last_name, rating, number_of_cancellations, email, email_verified, profile_pic_url, tac_version, emergency_contact, status, created_time, updated_time, usage_details) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	results, err := database.DB.Exec(
		query,
		userObject.Id,
		userObject.PhoneNumber,
		userObject.UserLevel,
		userObject.Gender,
		userObject.UserSubType,
		userObject.PaymentMethods,
		userObject.WalletAmount,
		userObject.FavoriteLocations,
		userObject.FirstName,
		userObject.LastName,
		userObject.Rating,
		userObject.NumberOfCancellations,
		userObject.Email,
		userObject.EmailVerified,
		userObject.ProfilePicUrl,
		userObject.TacVersion,
		userObject.EmergencyContact,
		userObject.Status,
		userObject.CreatedTime,
		userObject.UpdatedTime,
		userObject.UsageDetails,
	)

	if err != nil {
		c.JSON(helper.GetHTTPError("Failed to insert data: "+err.Error(), http.StatusInternalServerError, c.FullPath()))
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"data":    results,
	})

}

func GetAllCustomers(c *gin.Context) {

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

	// Fetch all customers from the database
	query := `
SELECT id, phone_number, user_level, gender, user_sub_type, payment_methods, wallet_amount, favorite_locations, 
first_name, last_name, rating, number_of_cancellations, email, email_verified, profile_pic_url, 
tac_version, emergency_contact, status, created_time, updated_time, usage_details 
FROM customer`

	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(helper.GetHTTPError("Failed to fetch customers: "+err.Error(), http.StatusInternalServerError, c.FullPath()))
		return
	}
	defer rows.Close()

	// Create a slice to hold the fetched customers
	var customers []models.User

	for rows.Next() {
		var customer models.User
		err := rows.Scan(
			&customer.Id,
			&customer.PhoneNumber,
			&customer.UserLevel,
			&customer.Gender,
			&customer.UserSubType,
			&customer.PaymentMethods,
			&customer.WalletAmount,
			&customer.FavoriteLocations,
			&customer.FirstName,
			&customer.LastName,
			&customer.Rating,
			&customer.NumberOfCancellations,
			&customer.Email,
			&customer.EmailVerified,
			&customer.ProfilePicUrl,
			&customer.TacVersion,
			&customer.EmergencyContact,
			&customer.Status,
			&customer.CreatedTime,
			&customer.UpdatedTime,
			&customer.UsageDetails,
		)
		if err != nil {
			c.JSON(helper.GetHTTPError("Failed to parse customer data: "+err.Error(), http.StatusInternalServerError, c.FullPath()))
			return
		}
		customers = append(customers, customer)
	}

	// Check if no customers were found
	if len(customers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No customers found"})
		return
	}

	// Return fetched customers
	c.JSON(http.StatusOK, gin.H{"data": customers})
}
