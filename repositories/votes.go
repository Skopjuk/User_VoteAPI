package repositories

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
	"userapi/models"
)

type VotesRepository struct {
	db *sqlx.DB
}

func NewVotesRepository(db *sqlx.DB) *VotesRepository {
	return &VotesRepository{db: db}
}

func (u *VotesRepository) AddVoteRecord(vote models.Votes) error {
	query := "INSERT INTO votes (user_id, rated_user_id, vote) values ($1, $2, $3)"
	_, err := u.db.Query(query, vote.UserId, vote.RatedUserId, vote.Vote)
	if err != nil {
		logrus.Errorf("vote wasn't inserted to DB: %s", err)
	}
	return err
}

func (u *VotesRepository) ChangeVote(vote models.Votes) (err error) {
	query := "UPDATE votes SET rated_user_id=$1, vote=$2 WHERE user_id=$3"
	_, err = u.db.Query(query, vote.RatedUserId, vote.Vote, vote.UserId)
	if err != nil {
		logrus.Errorf("problem with query while updating user: %s", err)
	}

	return err
}

func (u *VotesRepository) GetAllVotes() (votesList []models.Votes, err error) {
	query := "SELECT * FROM votes"
	err = u.db.Select(&votesList, query)
	if err != nil {
		logrus.Errorf("votes were not found: %s", err)
	}

	return votesList, err
}

func (u *VotesRepository) GetUsersRate(id int) (vote int, err error) {
	query := "SELECT vote FROM votes WHERE user_id=$1"
	err = u.db.Get(&vote, query, id)
	if err != nil {
		logrus.Errorf("rate record for user with id %d wasn't find: %s", id, err)
	}

	return vote, err
}

func (u *VotesRepository) DeleteVote(userId, ratedUserId int) error {
	query := "DELETE FROM votes WHERE user_id=$1 AND rated_user_id=$2"
	_, err := u.db.Query(query, userId, ratedUserId)
	if err != nil {
		logrus.Errorf("qwery for deleting vote can not be executed")
	}

	return err
}

func (u *VotesRepository) GetVoteByUserIds(userWhoVote, userForWhomVote int) (vote int, err error) {
	query := "SELECT vote FROM votes WHERE user_id=$1 AND rated_user_id=$2 LIMIT 1"
	err = u.db.Get(&vote, query, userWhoVote, userForWhomVote)
	if vote != 0 {
		newErr := fmt.Sprintf("user with id %d already voted for user with id %d", userWhoVote, userForWhomVote)
		logrus.Errorf(newErr)
		return vote, errors.New(newErr)
	}

	return vote, nil
}

func (u *VotesRepository) FindUsersVeryLastVote(voterId int) (updatedAt time.Time, err error) {
	query := "SELECT updated_at FROM votes WHERE user_id=$1 ORDER BY updated_at DESC LIMIT 1"
	err = u.db.Get(&updatedAt, query, voterId)
	if err != nil {
		logrus.Errorf("record for user with id %d wasn't find: %s", voterId, err)
	}

	return updatedAt, err
}
