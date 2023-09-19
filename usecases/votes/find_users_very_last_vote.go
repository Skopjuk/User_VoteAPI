package votes

import (
	"time"
	"userapi/usecases/user"
)

type FindLastVote struct {
	repository user.FindUsersVeryLastVote
}

func NewFindLastVote(repository user.FindUsersVeryLastVote) *FindLastVote {
	return &FindLastVote{repository: repository}
}

func (f *FindLastVote) Execute(voterId int) (time time.Time, err error) {
	return f.repository.FindUsersVeryLastVote(voterId)
}
