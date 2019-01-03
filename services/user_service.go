package services

import (
	"casbin-demo/db/mappers"
	"casbin-demo/models"
	"casbin-demo/utils"
	"time"
)

type UserService interface {
	DoRegiste(user *models.User) error
	DoLogin(user *models.User) (bool, error)
}

func NewUserService(userMapper mappers.UserMapper) UserService {
	return &userService{
		repo: userMapper,
	}
}

type userService struct {
	repo mappers.UserMapper
}

func (us *userService) DoRegiste(user *models.User) error {
	user.CreateTime = time.Now()
	user.Password = utils.AESEncrypt([]byte(user.Password))

	return us.repo.RegisteUser(user)
}

func (us *userService) DoLogin(user *models.User) (bool, error) {
	return us.repo.QueryByUsername(user)
}
