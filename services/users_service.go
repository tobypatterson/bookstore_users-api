package services

import (
	"github.com/tobypatterson/bookstore_users-api/domain/users"
	"github.com/tobypatterson/bookstore_users-api/utils/crypt_utils"
	"github.com/tobypatterson/bookstore_users-api/utils/date_utils"
	"github.com/tobypatterson/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	SearchUser(string) ([]users.User, *errors.RestErr)
}

// CreateUser will create a user
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDbFormat()
	user.Status = "active"
	user.Password = crypt_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser will delete a user
func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := users.User{Id: userId}

	return user.Delete()
}

// GetUser will return a user
func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("Invalid User ID")
	}

	result := &users.User{Id: userId}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateUser will update the provided user
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {

	current, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

// SearchUser will find users by a given status
func (s *usersService) SearchUser(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}

	return dao.FindUserByStatus(status)
}
