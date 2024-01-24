package adapters

import (
	"neptune/entities"
	"neptune/usecases"

	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) usecases.OrderRepository {
	return &GormOrderRepository{db: db}
}

// Save implements usecases.OrderRepository
func (*GormOrderRepository) Save(entities.Order) error {
	panic("unimplemented")
}
