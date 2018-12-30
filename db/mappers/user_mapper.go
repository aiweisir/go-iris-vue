package mappers

import (
	"casbin-demo/db"
	"casbin-demo/models"

	"github.com/jinzhu/gorm"
	"github.com/kataras/golog"
)

//type Query func(user models.User) bool

type UserMapper interface {
	QueryByUsername(username string) *models.User
	InsertOneUser(user *models.User) bool
}

func NewUserMapper() UserMapper {
	return &userMapper{
		db: db.MasterDB(),
	}
}

type userMapper struct {
	db *gorm.DB
}

func (m *userMapper) QueryByUsername(username string) *models.User {
	//// 读取
	//var product Product
	//db.First(&product, 1) // 查询id为1的product
	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	user := models.User{}
	golog.Info("models.User{} is nil")
	m.db.First(&user, "username=?", username)
	return &user
}

func (m *userMapper) InsertOneUser(user *models.User) bool {
	m.db.Create(&user)
	f := m.db.NewRecord(user)
	return f
}



