package services

import (
	"go-iris/web/db/mappers"
	"go-iris/web/models"
)

type (
	DemoService interface {
		AddOneProduct(demo *models.Demo) error
		GetOneProduct(demo *models.Demo) (bool, error)
	}

	demoService struct {
		repo mappers.DemoMapper
	}
)

func NewDemoService(demoMapper mappers.DemoMapper) DemoService {
	return &demoService{
		repo: demoMapper,
	}
}

func (ds *demoService) AddOneProduct(demo *models.Demo) error {
	return ds.repo.InsertOneDemo(demo)
}

func (ds *demoService) GetOneProduct(demo *models.Demo) (bool, error) {
	return ds.repo.QueryById(demo)
}
