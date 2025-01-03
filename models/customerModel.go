package models

type User struct {
	Id                    string `json:"id" binding:"required"`
	PhoneNumber           string `json:"phone_number" binding:"required"`
	UserLevel             string `json:"user_level" binding:"required"`
	Gender                string `json:"gender" binding:"required"`
	UserSubType           string `json:"user_sub_type" binding:"required"`
	PaymentMethods        string `json:"payment_methods" binding:"required"`
	WalletAmount          int    `json:"wallet_amount" binding:"required"`
	FavoriteLocations     string `json:"favorite_locations" binding:"required"`
	FirstName             string `json:"first_name" binding:"required"`
	LastName              string `json:"last_name" binding:"required"`
	Rating                int    `json:"rating" binding:"required"`
	NumberOfCancellations int    `json:"number_of_cancellations" binding:"required"`
	Email                 string `json:"email" binding:"required"`
	EmailVerified         int    `json:"email_verified" binding:"required"`
	ProfilePicUrl         string `json:"profile_pic_url" binding:"required"`
	TacVersion            string `json:"tac_version" binding:"required"`
	EmergencyContact      string `json:"emergency_contact" binding:"required"`
	Status                string `json:"status" binding:"required"`
	CreatedTime           string `json:"created_time" binding:"required"`
	UpdatedTime           string `json:"updated_time" binding:"required"`
	UsageDetails          string `json:"usage_details" binding:"required"`
}
