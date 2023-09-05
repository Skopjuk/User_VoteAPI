package models

import "time"

type User struct {
	Id        int        `db:"id"`
	Username  string     `db:"username"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Password  []byte     `db:"password"`
	Role      string     `db:"role"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
