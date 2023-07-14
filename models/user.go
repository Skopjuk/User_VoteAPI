package models

import "time"

type User struct {
	Id        int
	Username  string
	FirstName string
	LastName  string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
