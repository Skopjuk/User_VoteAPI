package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"userapi/usecases/rating"
)

func (r *RatingHandler) GetRatingByUserId(c echo.Context) error {
	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal error",
		})
	}

	newGetUserRatingById := rating.NewGetUserRating(r.container.Repository)
	userRating, err := newGetUserRatingById.Execute(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = c.JSON(http.StatusOK, map[string]interface{}{
		"rating": userRating,
	})

	if err != nil {
		logrus.Errorf("troubles with sending http status: %s", err)
	}

	return err
}