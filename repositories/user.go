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
		logrus.Errorf("user with id %s wasn't found", id)
	}

	return id, err
}

func (u *UsersRepository) FindUserByUsername(username string) (user models.User, err error) {
	query := "SELECT * FROM users WHERE username=$1 LIMIT 1"
	err = u.db.Get(&user, query, username)
	if err != nil {
		logrus.Errorf("user %s wasn't found", user)
	}

	return user, err
}

func (u *UsersRepository) UpdateUser(user models.User, id int) error {
	query := "UPDATE users SET username=$1, first_name=$2, last_name=$3 WHERE id=$4"
	_, err := u.db.Query(query, user.Username, user.FirstName, user.LastName, id)
	if err != nil {
		logrus.Errorf("query problem: %s", err)
	}

	return err
}

func (u *UsersRepository) GetAll(skip string) (usersList []models.User, err error) {
	query := "SELECT * FROM users LIMIT 10 OFFSET $1"
	err = u.db.Select(&usersList, query, skip)

	if err != nil {
		logrus.Errorf("users not found %s")
	}

	return usersList, err
}

func (u *UsersRepository) GetUserById(id int) (user models.User, err error) {
	query := "SELECT * FROM users WHERE id=$1"
	err = u.db.Get(&user, query, id)
	if err != nil {
		logrus.Errorf("user with id %s wasn't found, with error: %s", id, err)
	}

	return user, err
}

func (u *UsersRepository) ChangeUsersPassword(id int, password string) error {
	query := "UPDATE users SET password=$1 WHERE id=$2"
	_, err := u.db.Query(query, password, id)
	if err != nil {
		logrus.Errorf("query for deleting password can not execute")
	}

	return err
}

func (u *UsersRepository) CountUsers() (numberOfUsers int, err error) {
	query := "SELECT COUNT(*) FROM users"
	err = u.db.Get(&numberOfUsers, query)
	if err != nil {
		logrus.Errorf("query for counting users returned error: %s", err)
	}

	return numberOfUsers, err
}
