package services

import (
	"casbin-demo/db/mappers"
	"casbin-demo/models"
)

type UserService interface {
	DoRegiste(user *models.User) bool
	DoLogin(username string) *models.User
}

func NewUserService(userMapper mappers.UserMapper) UserService {
	return &userService{
		repo: userMapper,
	}
}

type userService struct {
	repo mappers.UserMapper
}

func (us *userService) DoRegiste(user *models.User) bool {
	return us.repo.InsertOneUser(user)
}
func (us *userService) DoLogin(username string) *models.User {
	return us.repo.QueryByUsername(username)
}
