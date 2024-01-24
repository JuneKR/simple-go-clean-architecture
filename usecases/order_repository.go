package usecases

import (
	"neptune/entities"
)

type OrderRepository interface {
	Save(entities.Order) error
}
