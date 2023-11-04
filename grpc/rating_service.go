package grpc

import (
	"context"
	"github.com/sirupsen/logrus"
	"strconv"
	"userapi/container"
	"userapi/usecases/rating"
)

type RateServer struct {
	Container *container.Container
}

func (r RateServer) GetRatingByUserId(c context.Context, request *GetRatingByUserIdRequest) (responce *GetRatingByUserIdResponce, err error) {
	redisKey := "rating_of_user_with_id_" + strconv.Itoa(int(request.Id))

	ratingRedis, err := r.Container.RedisDb.Get(c, redisKey).Result()
	if err != nil {
		logrus.Errorf("error while getting data from redis: %s", err)
	}

	if ratingRedis != "" {
		logrus.Info("data about users list exists in redis")

		ratingInt, err := strconv.Atoi(ratingRedis)
		if err != nil {
			logrus.Errorf("error while converting rating to int")
			return nil, err
		}

		responce = &GetRatingByUserIdResponce{Rating: int32(ratingInt)}
		return responce, nil
	}
	newGetUserRatingById := rating.NewGetUserRating(r.Container.RatingRepository)
	userRating, err := newGetUserRatingById.Execute(int(request.Id))
	if err != nil {
		return nil, err
	}

	r.Container.RedisDb.Set(c, redisKey, userRating, r.Container.Config.ExpTime)

	responce = &GetRatingByUserIdResponce{Rating: int32(userRating)}

	return responce, nil
}

func (r RateServer) mustEmbedUnimplementedRatingServer() {}
