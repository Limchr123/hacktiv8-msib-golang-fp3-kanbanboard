package dto

import "time"

type NewUserRequest struct {
	FullName string `json:"full_name" valid:"required~full_name cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
}

type LoginRequest struct {
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
}

type NewUser struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type NewUserResponse struct {
	Status int     `json:"status"`
	Data   NewUser `json:"data"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Status int           `json:"status"`
	Data   TokenResponse `json:"data"`
}
