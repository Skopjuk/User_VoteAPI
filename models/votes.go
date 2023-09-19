package models

import "time"

type Votes struct {
	Id          int       `db:"id"`
	UserId      int       `db:"user_id"`
	RatedUserId int       `db:"rated_user_id"`
	Vote        int       `db:"vote"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
