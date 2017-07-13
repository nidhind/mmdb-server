package data_models

// Model of Login request body
type LoginBody struct {
	UserId string `json:"userid" binding:"required"`
	Password string `json:"password" binding:"required"`
}