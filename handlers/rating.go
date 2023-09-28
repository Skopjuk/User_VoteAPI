package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"userapi/usecases/rating"
)

func (r *RatingHandler) GetRatingByUserId(c echo.Context) error {
	idInt, err := getIdFromEndpoint(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Internal error",
		})
	}

	redisKey := "rating_of_user_with_id_" + strconv.Itoa(idInt)

	ratingRedis, err := r.container.RedisDb.Get(c.Request().Context(), redisKey).Result()

	if ratingRedis != "" {
		logrus.Info("data about users list exists in redis")
		return c.JSON(http.StatusOK, map[string]interface{}{
			"users_rating": ratingRedis,
		})
	}
	newGetUserRatingById := rating.NewGetUserRating(r.container.RatingRepository)
	userRating, err := newGetUserRatingById.Execute(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	data, err := json.Marshal(userRating)
	if err != nil {
		logrus.Errorf("error while marshaling users list")
		return err
	}

	r.container.RedisDb.Set(c.Request().Context(), redisKey, data, r.container.Config.ExpTime)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"rating": userRating,
	})
}
