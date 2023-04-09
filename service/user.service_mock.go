package service

import (
	"errors"

	"github.com/mnfirdauss/go-gorm/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (um *UserRepositoryMock) CreateUser(user *model.User) error {
	// argument := um.Mock.Called(user)
	if user == nil {
		return errors.New("error")
	} else {
		return nil
	}
}
