package models

import (
	"encoding/json"
	"time"
)

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

func (m User) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m User) UnmarshalBinary(toUnmarshal []byte) error {
	res := User{}
	return json.Unmarshal(toUnmarshal, &res)
}
