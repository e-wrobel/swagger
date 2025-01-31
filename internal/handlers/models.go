package handlers

import "time"

type CreateUserRequest struct {
	User
	Address
}

type User struct {
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Age     time.Time `json:"age"`
}

type Address struct {
	Street string `json:"street"`
}

type CreateUserResponse struct {
	Data map[string]string `json:"data"`
}
