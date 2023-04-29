package model

type LoginRequest struct {
	// UserInfo could be an email, username or phone number
	UserInfo string
	Password string
}
