package user

import "fmt"

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
	if err := validatePassword(attributes); err != nil {
		return err
	}

	return a.repository.ChangeUsersPassword(id, string(PasswordHashing(attributes.Password)))
}

func validatePassword(attributes ChangePasswordAttributes) error {
	if len(attributes.Password) < 6 {
		return fmt.Errorf("password is too short")
	} else if len(attributes.Password) > 100 {
		return fmt.Errorf("password is too long")
	}

	return nil
}