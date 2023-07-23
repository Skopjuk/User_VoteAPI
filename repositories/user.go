package repositories

import (
	"github.com/jmoiron/sqlx"
	//	"userapi"
	"userapi/models"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (u *UsersRepository) InsertUser(user models.User) (id int, err error) {
	query := "INSERT INTO users (username, password, first_name, last_name) values ($1, $2, $3, $4) RETURNING id"
	row := u.db.QueryRow(query, user.Username, user.Password, user.FirstName, user.LastName)
	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
