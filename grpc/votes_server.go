package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
	"userapi/container"
	"userapi/handlers"
	"userapi/repositories"
	"userapi/usecases/rating"
	"userapi/usecases/votes"
)

type ServerVotes struct {
	Container *container.Container
}

func (s ServerVotes) AddVote(c context.Context, request *AddVoteRequest) (*emptypb.Empty, error) {
	err := s.CheckIfUserCanVote(int(request.Id), int(request.RatedUserId), c)
	if err != nil {
		return new(emptypb.Empty), err
	}

	err = s.vote(int(request.Id), int(request.RatedUserId), int(request.Vote))
	if err != nil {
		return new(emptypb.Empty), err
	}

	newCreateRating := rating.NewCreateUserRating(s.Container.RatingRepository)
	newUpdateRating := rating.NewUpdateUsersRating(s.Container.RatingRepository)
	newGetUserRating := rating.NewGetUserRating(s.Container.RatingRepository)

	newCreatingOrUpdatingRating := rating.NewCreateOrUpdateRating(newCreateRating, newUpdateRating, newGetUserRating)
	_, err = newCreatingOrUpdatingRating.Execute(int(request.RatedUserId), int(request.Vote))
	if err != nil {
		return new(emptypb.Empty), err
	}

	return new(emptypb.Empty), nil
}

func (s ServerVotes) ChangeVote(ctx context.Context, request *ChangeVoteRequest) (*emptypb.Empty, error) {
	changeRateAttributes := votes.ChangeRateAttributes{
		UserId:      int(request.UserId),
		RatedUserId: int(request.RatedUserId),
		Vote:        int(request.Vote),
	}
	newChangeVote := votes.NewChangeVote(s.Container.VotesRepository)
	err := newChangeVote.Execute(changeRateAttributes)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
	}

	return new(emptypb.Empty), err
}

func (s ServerVotes) GetAllVotes(c context.Context, in *emptypb.Empty) (responce *GetAllVotesResponce, err error) {
	var votesList []*Vote
	newGetVotes := votes.NewGetListOfVotes(s.Container.VotesRepository)

	votes, err := newGetVotes.Execute()
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return nil, err
	}

	for _, i := range votes {
		vote := Vote{
			Id:          int32(i.Id),
			UsersId:     int32(i.UserId),
			RatedUserId: int32(i.RatedUserId),
			Vote:        int32(i.Vote),
		}

		votesList = append(votesList, &vote)
	}

	responce = &GetAllVotesResponce{
		Vote: votesList,
	}

	return responce, err
}

func (s ServerVotes) DeleteVote(c context.Context, in *DeleteVoteRequest) (*emptypb.Empty, error) {
	newGetVoteByUsersId := votes.NewGetVoteByUsersId(s.Container.VotesRepository)
	vote, err := newGetVoteByUsersId.Execute(int(in.UserId), int(in.RatedUserId))
	if err.Error() != fmt.Sprintf("user with id %d already voted for user with id %d", in.UserId, in.RatedUserId) && err != nil {
		return new(emptypb.Empty), err
	}

	newGetUserRating := rating.NewGetUserRating(s.Container.RatingRepository)
	userRating, err := newGetUserRating.Execute(int(in.RatedUserId))
	newRating := userRating - vote
	err = handlers.UpdateUsersRating(int(in.RatedUserId), newRating, *s.Container)
	if err != nil {
		return new(emptypb.Empty), err
	}

	newDeleteVote := votes.NewDeleteUsersVote(s.Container.VotesRepository)
	err = newDeleteVote.Execute(int(in.UserId), int(in.RatedUserId))
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return new(emptypb.Empty), err
	}

	return new(emptypb.Empty), err
}

func (s ServerVotes) mustEmbedUnimplementedVoteServiceServer() {}

func (s *ServerVotes) CheckIfUserCanVote(userId, RatedUserId int, c context.Context) error {
	newCheckIfUserAlreadyVoted := votes.NewGetVoteByUsersId(s.Container.VotesRepository)
	_, err := newCheckIfUserAlreadyVoted.Execute(userId, RatedUserId)
	if err.Error() != "sql: no rows in result set" && err != nil {
		return err
	}

	newOneHourCheck := votes.NewFindLastVote(s.Container.VotesRepository)
	updatedAt, err := newOneHourCheck.Execute(userId)
	if err != nil {
		logrus.Errorf("error while looking for very last users vote: %s", err)
	}

	timeDiff := time.Now().Hour() - updatedAt.Hour()
	if timeDiff < 1 {
		return errors.New("you can vote maximum once per hour")
	}

	return nil
}

func (s *ServerVotes) vote(userId, ratedUserId, vote int) error {
	userRepository := repositories.NewVotesRepository(s.Container.DB)
	newVote := votes.NewVote(userRepository)

	params := votes.NewVoteAttributes{
		UserId:      userId,
		RatedUserId: ratedUserId,
		Vote:        vote,
	}
	err := newVote.Execute(params)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
	}

	return err
}
