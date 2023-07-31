package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

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

func (u *UsersRepository) FindUserByUsername(username string) (user models.User, err error) {
	query := "SELECT * FROM users WHERE username=$1 LIMIT 1"
	err = u.db.Get(&user, query, username)
	if err != nil {
		logrus.Errorf("user %s wasn't found", user)
	}

	return user, err
}

func (u *UsersRepository) UpdateUser(user models.User) error {
	query := "UPDATE users SET username=$1, first_name=$3, last_name=$4 values ($1, $2, $3, $4)"
	_, err := u.db.Query(query, user.Username, user.FirstName, user.LastName)
	if err != nil {
		logrus.Errorf("query problem: %s", err)
	}

	return err
}

func (u *UsersRepository) GetAll() (usersList []models.User, err error) {
	query := "SELECT * FROM users"
	err = u.db.Select(&usersList, query)

	if err != nil {
		logrus.Errorf("users not found %s")
	}

	return usersList, err
}
