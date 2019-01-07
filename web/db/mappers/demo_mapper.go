package mappers

import (
	"go-iris/web/db"
	"go-iris/web/models"

	"github.com/go-xorm/xorm"
)

type (
	DemoMapper interface {
		InsertOneDemo(demo *models.Demo) error
		QueryById(demo *models.Demo) (bool, error)
	}

	demoMapper struct {
		db *xorm.Engine
	}
)

func NewDemoMapper() DemoMapper {
	return &demoMapper{
		db: db.MasterEngine(),
	}
}

func (dm *demoMapper) InsertOneDemo(demo *models.Demo) error {
	_, err := dm.db.Insert(demo)
	return err
}

func (dm *demoMapper) QueryById(demo *models.Demo) (bool, error) {
	return dm.db.Get(demo)
}
