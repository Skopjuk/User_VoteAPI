package models

import "time"

type Rating struct {
	Id        string    `db:"id"`
	UserId    int       `db:"user_id"`
	Rating    int       `db:"rating"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
