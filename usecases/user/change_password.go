package user

type ChangePassword struct {
	//	repository          AuthenticateUser
	repository ChangeUsersPassword
}

func NewChangePassword(repository ChangeUsersPassword) *ChangePassword {
	return &ChangePassword{repository: repository}
}

type ChangePasswordAttributes struct {
	Password string
}

func (a *ChangePassword) Execute(id int, attributes ChangePasswordAttributes) error {
	//authenticated := a.repository.AuthenticateUser(attributes.Username, attributes.Password)
	//if !authenticated {
	//	logrus.Error("user is not authenticated")
	//	return false, fmt.Errorf("user is not authenticated")
	//}

	return a.repository.ChangeUsersPassword(id, string(PasswordHashing(attributes.Password)))
}
