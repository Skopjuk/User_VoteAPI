package repositories

import (
	"errors"
	"fmt"
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
	query := "INSERT INTO users (username, password, first_name, last_name, role) values ($1, $2, $3, $4, $5) RETURNING id"
	row := u.db.QueryRow(query, user.Username, user.Password, user.FirstName, user.LastName, user.Role)
	if err = row.Scan(&id); err != nil {
		logrus.Errorf("user with id %s wasn't found", id)
	}

	return id, err
}

func (u *UsersRepository) AddVoteRecord(vote models.Rate) error {
	query := "INSERT INTO rates (user_id, rated_user_id, username_who_votes, username_for_whom_votes, rate) values ($1, $2, $3, $4, $5)"
	_, err := u.db.Query(query, vote.UserId, vote.RatedUserId, vote.UsernameWhoVotes, vote.UsernameForWhomVotes, vote.Rate)
	if err != nil {
		logrus.Errorf("vote wasn't inserted to DB: %s", err)
	}
	return err
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

func (u *UsersRepository) ChangeVote(vote models.Rate, id int) (err error) {
	query := "UPDATE rates SET user_id=$1, rated_user_id=$2, username_who_votes=$3, username_for_whom_votes=$4, rate=$5 WHERE id=$6"
	_, err = u.db.Query(query, vote.UserId, vote.RatedUserId, vote.UsernameWhoVotes, vote.UsernameForWhomVotes, vote.Rate, id)
	if err != nil {
		logrus.Errorf("problem with query while updating user: %s", err)
	}

	return err
}

func (u *UsersRepository) GetAll(skip string, paginationLimit string) (usersList []models.User, err error) {
	query := "SELECT * FROM users LIMIT $1 OFFSET $2"
	err = u.db.Select(&usersList, query, paginationLimit, skip)

	if err != nil {
		logrus.Errorf("users not found %s")
	}

	return usersList, err
}

func (u *UsersRepository) GetAllVotes() (votesList []models.Rate, err error) {
	query := "SELECT * FROM rates"
	err = u.db.Select(&votesList, query)
	if err != nil {
		logrus.Errorf("votes were not found: %s", err)
	}

	return votesList, err
}

func (u *UsersRepository) GetUserById(id int) (user models.User, err error) {
	query := "SELECT * FROM users WHERE id=$1"
	err = u.db.Get(&user, query, id)
	if err != nil {
		logrus.Errorf("user with id %s wasn't found, with error: %s", id, err)
	}

	return user, err
}

func (u *UsersRepository) GetUsersRate(id int) (rate int, err error) {
	query := "SELECT rate FROM rates WHERE user_id=$1"
	err = u.db.Get(&rate, query, id)
	if err != nil {
		logrus.Errorf("rate record for user with id %d wasn't find: %s", id, err)
	}

	return rate, err
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

func (u *UsersRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := u.db.Query(query, id)
	if err != nil {
		logrus.Errorf("query for deleting user can not be executed")
	}

	return err
}

func (u *UsersRepository) DeleteVote(id int) error {
	query := "DELETE FROM rates WHERE id=$1"
	_, err := u.db.Query(query, id)
	if err != nil {
		logrus.Errorf("qwery for deleting vote can not be executed")
	}

	return err
}

func (u *UsersRepository) CheckIfUserVotedForSomeUser(userWhoVote, userForWhomVote int) (err error) {
	var row models.Rate
	query := "SELECT * FROM rates WHERE user_id=$1 AND rated_user_id=$2 LIMIT 1"
	err = u.db.Get(&row, query, userWhoVote, userForWhomVote)
	if err == nil {
		newErr := fmt.Sprintf("user with id %d already voted for user with id %d", userWhoVote, userForWhomVote)
		logrus.Errorf(newErr)
		return errors.New(newErr)
	}

	return nil
}
