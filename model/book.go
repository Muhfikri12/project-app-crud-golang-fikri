package model

import "time"

type Books struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Publish string `json:"publish"`
	IsBorrowed bool `json:"isBorrowed"`
	Created_at time.Time `json:"created_at"`
}


