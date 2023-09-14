package models

import "time"

type Rate struct {
	Id                   int       `db:"id"`
	UserId               int       `db:"user_id"`
	RatedUserId          int       `db:"rated_user_id"`
	UsernameWhoVotes     string    `db:"username_who_votes"`
	UsernameForWhomVotes string    `db:"username_for_whom_votes"`
	Rate                 int       `db:"rate"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}
