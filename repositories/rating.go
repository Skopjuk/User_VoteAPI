package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
	"userapi/models"
)

type RatingRepository struct {
	db *sqlx.DB
}

func NewRatingRepository(db *sqlx.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

func (u *RatingRepository) AddUserRating(rating models.Rating) error {
	query := "INSERT INTO ratings (user_id, rating) values ($1, $2)"
	_, err := u.db.Query(query, rating.UserId, rating.Rating)

	return err
}

func (u *RatingRepository) GetRatingByUserId(id int) (rating int, err error) {
	query := "SELECT rating FROM ratings WHERE user_id=$1"

	err = u.db.Get(&rating, query, id)
	if err != nil {
		logrus.Errorf("users rating wasn't find")
	}

	return rating, err
}

func (u *RatingRepository) UpdateUserRating(rating, id int) error {
	query := "UPDATE ratings SET rating=$1, updated_at=$2 WHERE user_id=$3"

	_, err := u.db.Query(query, rating, time.Now(), id)
	if err != nil {
		logrus.Errorf("users rating wasn't updated")
	}

	return err
}

func (u *RatingRepository) DeleteUserRating(userId int) error {
	query := "DELETE FROM ratings WHERE user_id=$1"

	_, err := u.db.Query(query, userId)
	if err != nil {
		logrus.Errorf("users rating wasn't deleted")
	}

	return err
}
