package model

import "time"

type Borrower struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Book Books `json:"book"`	
	Status bool `json:"status"`
	Created_at time.Time `json:"created_at"`
}

type User struct {
	Username string
	Password string
}

var ValidUser = User{
	Username: "admin",
	Password: "password",
}