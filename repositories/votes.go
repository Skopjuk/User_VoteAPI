package repositories

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"userapi/models"
)

func (u *UsersRepository) AddVoteRecord(vote models.Votes) error {
	query := "INSERT INTO votes (user_id, rated_user_id, vote) values ($1, $2, $3)"
	_, err := u.db.Query(query, vote.UserId, vote.RatedUserId, vote.Vote)
	if err != nil {
		logrus.Errorf("vote wasn't inserted to DB: %s", err)
	}
	return err
}

func (u *UsersRepository) ChangeVote(vote models.Votes, id int) (err error) {
	query := "UPDATE votes SET user_id=$1, rated_user_id=$2, vote=$3 WHERE id=$4"
	_, err = u.db.Query(query, vote.UserId, vote.RatedUserId, vote.Vote, id)
	if err != nil {
		logrus.Errorf("problem with query while updating user: %s", err)
	}

	return err
}

func (u *UsersRepository) GetAllVotes() (votesList []models.Votes, err error) {
	query := "SELECT * FROM votes"
	err = u.db.Select(&votesList, query)
	if err != nil {
		logrus.Errorf("votes were not found: %s", err)
	}

	return votesList, err
}

func (u *UsersRepository) GetUsersRate(id int) (vote int, err error) {
	query := "SELECT vote FROM votes WHERE user_id=$1"
	err = u.db.Get(&vote, query, id)
	if err != nil {
		logrus.Errorf("rate record for user with id %d wasn't find: %s", id, err)
	}

	return vote, err
}

func (u *UsersRepository) DeleteVote(userId, ratedUserId int) error {
	query := "DELETE FROM votes WHERE user_id=$1 AND rated_user_id=$2"
	_, err := u.db.Query(query, userId, ratedUserId)
	if err != nil {
		logrus.Errorf("qwery for deleting vote can not be executed")
	}

	return err
}

func (u *UsersRepository) GetVoteByUserIds(userWhoVote, userForWhomVote int) (vote int, err error) {
	query := "SELECT vote FROM votes WHERE user_id=$1 AND rated_user_id=$2 LIMIT 1"
	err = u.db.Get(&vote, query, userWhoVote, userForWhomVote)
	if err == nil {
		newErr := fmt.Sprintf("user with id %d already voted for user with id %d", userWhoVote, userForWhomVote)
		logrus.Errorf(newErr)
		return vote, errors.New(newErr)
	}

	return vote, nil
}

func (u *UsersRepository) FindUsersVeryLastVote(voterId int) (updatedAt time.Time, err error) {
	query := "SELECT updated_at FROM votes WHERE user_id=$1 ORDER BY updated_at DESC LIMIT 1"
	err = u.db.Get(&updatedAt, query, voterId)
	if err != nil {
		logrus.Errorf("record for user with id %d wasn't find: %s", voterId, err)
	}

	return updatedAt, err
}
