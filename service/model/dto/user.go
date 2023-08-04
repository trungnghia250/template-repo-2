package dto

import "time"

type GetUserByIDRequest struct {
	ID string `query:"id"`
}

type GetUserByIDResponse struct {
	ID          string    `json:"id"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type CreateUserRequest struct {
	ID          string    `json:"id"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type UpdateUserRequest struct {
	ID          string    `json:"id,omitempty"`
	UserName    string    `json:"user_name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	DateOfBirth time.Time `json:"date_of_birth,omitempty"`
}

type DeleteUserRequest struct {
	ID string `query:"id"`
}
