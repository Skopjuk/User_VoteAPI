package grpc

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"userapi/container"
	"userapi/repositories"
	"userapi/usecases/user"
)

const paginationLimit = "10"

type Server struct {
	Container *container.Container
}

func (s Server) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {

	usersRepository := repositories.NewUsersRepository(s.Container.DB)
	newProfile := user.NewCreateProfile(usersRepository)

	params := user.NewUserAttributes{
		Username:  request.Username,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Role:      request.Role,
		Password:  request.Password,
	}

	id, err := newProfile.Execute(params)
	if err != nil {
		return nil, err
	}

	response := CreateUserResponse{
		Id: int64(id),
	}

	return &response, nil
}

func (s Server) FindUserByUsername(context.Context, *FindUserByUsernameRequest) (*FindUserByUsernameResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) UpdateUser(c context.Context, request *UpdateUserRequest) (*UpdateUserResponse, error) {
	input := user.UpdateUserAttributes{
		Username:  request.Username,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Role:      request.Role,
	}
	newUpdateProfile := user.NewChangeProfile(s.Container.UsersRepository)
	err := newUpdateProfile.Execute(input, int(request.Id))
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return nil, err
	}

	response := UpdateUserResponse{Username: request.Username}

	return &response, nil
}

func (s Server) GetAll(c context.Context, request *GetAllRequest) (*GetAllResponce, error) {
	var usersList []*User
	pageNum := request.Page
	var response *GetAllResponce
	skip := strconv.Itoa(int((pageNum - 1) * 10))

	newGetUsers := user.NewGetAllUsers(s.Container.UsersRepository)

	users, err := newGetUsers.Execute(skip, paginationLimit)
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return nil, err
	}

	for _, i := range users {
		user := User{
			Id:        int64(i.Id),
			Username:  i.Username,
			FirstName: i.FirstName,
			LastName:  i.LastName,
			Password:  i.Password,
			Role:      i.Role,
			StartedAt: timestamppb.New(i.CreatedAt),
			UpdatedAt: timestamppb.New(i.UpdatedAt),
		}
		usersList = append(usersList, &user)
	}

	response = &GetAllResponce{
		Users: usersList,
	}

	return response, nil
}

func (s Server) GetUserById(c context.Context, request *GetUserByIdRequest) (*GetUserByIdResponce, error) {
	var response User
	userById, err := s.Container.RedisDb.Get(c, string(request.Id)).Result()
	if err != nil {
		logrus.Errorf("error while getting data from redis")
	}

	if userById != "" {
		logrus.Info("data about this user exists in redis")
		if err := json.Unmarshal([]byte(userById), &response); err != nil {
			logrus.Errorf("failed to bind req body: %s", err)
			return nil, err
		}
		return &GetUserByIdResponce{User: &response}, nil
	}

	logrus.Info("in redis no data about this user. Request to Postrgres")

	newGetUserById := user.NewGetUserByID(s.Container.UsersRepository)
	foundUser, err := newGetUserById.Execute(int(request.Id))
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return nil, err
	}

	keyForRedis := "user_by_id_" + strconv.Itoa(int(request.Id))
	s.Container.RedisDb.Set(c, keyForRedis, foundUser, s.Container.Config.ExpTime)

	response = User{
		Id:        int64(foundUser.Id),
		Username:  foundUser.Username,
		FirstName: foundUser.FirstName,
		LastName:  foundUser.LastName,
		Password:  foundUser.Password,
		Role:      foundUser.Role,
	}

	return &GetUserByIdResponce{User: &response}, nil
}

func (s Server) ChangeUsersPassword(c context.Context, request *ChangeUsersPasswordRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) CountUsers(c context.Context, empty *emptypb.Empty) (*CountUsersResponse, error) {
	var responce CountUsersResponse
	usersNumRedis, err := s.Container.RedisDb.Get(c, "amount_of_users").Result()

	if usersNumRedis != "" {
		logrus.Info("data about amount of users exists in redis")

		usersNumRedisStr, err := strconv.Atoi(usersNumRedis)
		if err != nil {
			logrus.Errorf("error while converting amount of users to string: %s", err)
			return nil, err
		}

		responce = CountUsersResponse{
			NumberOfUsers: int32(usersNumRedisStr),
		}

		return &responce, nil
	}

	logrus.Info("in redis no data about amount of users. Request to Postrgres")
	newGetUserById := user.NewCountAllUsers(s.Container.UsersRepository)
	numOfUsers, err := newGetUserById.Execute()
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return nil, err
	}

	s.Container.RedisDb.Set(c, "amount_of_users", numOfUsers, s.Container.Config.ExpTime)

	responce = CountUsersResponse{
		NumberOfUsers: int32(numOfUsers),
	}

	logrus.Info(&responce)

	return &responce, nil
}

func (s Server) DeleteUser(c context.Context, request *DeleteUserRequest) (*emptypb.Empty, error) {
	newGetUserById := user.NewGetUserByID(s.Container.UsersRepository)
	_, err := newGetUserById.Execute(int(request.Id))
	if err != nil {
		logrus.Errorf("problem wile inserting user: %s", err)
		return new(emptypb.Empty), err
	}

	newDeleteUser := user.NewDeleteProfile(s.Container.UsersRepository)
	err = newDeleteUser.Execute(int(request.Id))
	if err != nil {
		logrus.Errorf("can not execute usecase: %s", err)
		return new(emptypb.Empty), err
	}
	return new(emptypb.Empty), nil
}

func (s Server) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}
