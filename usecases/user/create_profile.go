package user

import (
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"userapi/models"
)

type CreateProfile struct {
	repository InsertUser
}

type NewUserAttributes struct {
	Username  string
	FirstName string
	LastName  string
	Password  string
	Role      string
}

func NewCreateProfile(repository InsertUser) *CreateProfile {
	return &CreateProfile{repository: repository}
}

func (c *CreateProfile) Execute(attributes NewUserAttributes) (id int, err error) {
	err = parametersValidation(attributes)
	if err != nil {
		logrus.Errorf("error while creating user: %s", err)
		return 0, err
	}

	id, err = c.repository.InsertUser(models.User{
		Username:  attributes.Username,
		FirstName: attributes.FirstName,
		LastName:  attributes.LastName,
		Role:      attributes.Role,
		Password:  PasswordHashing(attributes.Password),
	})

	return id, err
}

func PasswordHashing(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		logrus.Error("Password wasn't hashed")
	}
	return hashedPassword
}

func parametersValidation(attributes NewUserAttributes) error {
	if len(attributes.FirstName) < 2 {
		return errors.New("first name is too short")
	} else if len(attributes.FirstName) > 50 {
		return errors.New("first name is too long")
	} else if len(attributes.LastName) < 2 {
		return errors.New("first name is too short")
	} else if len(attributes.LastName) > 50 {
		return errors.New("last name is too long")
	} else if attributes.Role != "user" && attributes.Role != "moderator" && attributes.Role != "admin" {
		return errors.New("user role is not valid")
	} else if len(attributes.Password) < 6 {
		return errors.New("password is too short")
	} else if len(attributes.Username) < 3 {
		return errors.New("username is too short")
	} else if len(attributes.Username) > 30 {
		return errors.New("username is too long")
	}
	return nil
}
