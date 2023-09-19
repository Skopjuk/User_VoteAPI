package repositories

import (
	"github.com/sirupsen/logrus"
	"userapi/models"
)

func (u *UsersRepository) AddUserRating(rating models.Rating) error {
	query := "INSERT INTO ratings (user_id, rating) values ($1, $2)"
	_, err := u.db.Query(query, rating.UserId, rating.Rating)

	return err
}

func (u *UsersRepository) GetRatingByUserId(id int) (rating int, err error) {
	query := "SELECT rating FROM ratings WHERE user_id=$1"

	err = u.db.Get(&rating, query, id)
	if err != nil {
		logrus.Errorf("users rating wasn't find")
	}

	return rating, err
}

func (u *UsersRepository) UpdateUserRating(rating, id int) error {
	query := "UPDATE ratings SET rating=$1 WHERE user_id=$2"

	_, err := u.db.Query(query, rating, id)
	if err != nil {
		logrus.Errorf("users rating wasn't updated")
	}

	return err
}

func (u *UsersRepository) DeleteUserRating(id int) error {
	query := "DELETE FROM ratings WHERE id=$1"

	_, err := u.db.Query(query, id)
	if err != nil {
		logrus.Errorf("users rating wasn't deleted")
	}

	return err
}
