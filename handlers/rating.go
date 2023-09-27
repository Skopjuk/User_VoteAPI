package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
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

	ratingRedis, err := r.container.RedisDb.Client.Get(r.container.Context, redisKey).Result()
	if err == redis.Nil {
		newGetUserRatingById := rating.NewGetUserRating(r.container.RatingRepository)
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

		data, err := json.Marshal(userRating)
		if err != nil {
			logrus.Errorf("error while marshaling users list")
		}

		status := r.container.RedisDb.Client.Set(r.container.Context, redisKey, data, 0)
		logrus.Info(status)
		_, err = r.container.RedisDb.Client.Expire(r.container.Context, redisKey, 1*time.Minute).Result()
		if err != nil {
			logrus.Errorf("error while set expirational period")
		}
	} else if err != nil {
		logrus.Errorf("error while attempt to recive data about users rating")
	} else {
		logrus.Info("data about users list exists in redis")
		err = c.JSON(http.StatusOK, map[string]interface{}{
			"users_rating": ratingRedis,
		})

		if err != nil {
			logrus.Errorf("troubles with sending http status: %s", err)
		}
	}

	return err
}
