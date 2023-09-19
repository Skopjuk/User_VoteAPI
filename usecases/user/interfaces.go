package user

import (
	"userapi/models"
)

type InsertUser interface {
	InsertUser(user models.User) (id int, err error)
}

type FindUserByUsername interface {
	FindUserByUsername(username string) (models.User, error)
}

type DeleteUser interface {
	DeleteUser(id int) error
}

type UpdateUser interface {
	UpdateUser(user models.User, id int) error
}

type AuthenticateUser interface {
	AuthenticateUser(username string, password string) bool
}

type ChangeUsersPassword interface {
	ChangeUsersPassword(id int, password string) error
}

type GetAll interface {
	GetAll(skip string, paginationLimit string) (usersList []models.User, err error)
}

type GetUserById interface {
	GetUserById(id int) (user models.User, err error)
}

type CountUsers interface {
	CountUsers() (numberOfUsers int, err error)
}

type CheckIfUserExists interface {
	CheckIfUserExist(id int) error
}

type AddVoteRecord interface {
	AddVoteRecord(vote models.Votes) error
}

type GetAllVotes interface {
	GetAllVotes() (votesList []models.Votes, err error)
}

type GetUsersRate interface {
	GetUsersRate(id int) (rate int, err error)
}

type CheckIfUserVotedForSomeUser interface {
	CheckIfUserVotedForSomeUser(userWhoVote, userForWhomVote int) (err error)
}

type ChangeVote interface {
	ChangeVote(vote models.Votes, id int) (err error)
}

type DeleteVote interface {
	DeleteVote(id int) (err error)
}
