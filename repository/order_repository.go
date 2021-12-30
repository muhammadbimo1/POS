package repository

import (
	"log"
	"table_management/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOne(order entity.CustomerOrder) (*entity.CustomerOrder, error)
}

type orderRepository struct {
	db *gorm.DB
}

func (o *orderRepository) CreateOne(order entity.CustomerOrder) (*entity.CustomerOrder, error) {
	err := o.db.Create(&order).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &order, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}
