package services

import (
	"go-iris/web/db/mappers"
	"go-iris/web/models"
	"go-iris/utils"
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
