package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"userapi/models"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) InsertUser(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO %s (username, password, first_name, last_name) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := u.db.QueryRow(query, user.Username, user.Password, user.FirstName, user.LastName)
	if err := row.Scan(query); err != nil {
		return err
	}
	return nil
}
