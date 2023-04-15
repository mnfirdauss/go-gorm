package service

import (
	"github.com/mnfirdauss/go-gorm/model"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (um *UserRepositoryMock) CreateUser(user *model.User) error {
	args := um.Mock.Called(user)
	return args.Error(0)
}
