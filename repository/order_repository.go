package repository

import (
	"log"
	"table_management/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOne(order entity.CustomerOrder) (*entity.CustomerOrder, error)
	GetSum(id string) (int, error)
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

func (o *orderRepository) GetSum(id string) (int, error) {
	var grandtotal int
	var customer entity.CustomerOrder
	result := o.db.Preload("OrderDetails").First(&customer, "id=?", id)
	for i := 0; i < len(customer.OrderDetails); i++ {
		grandtotal += customer.OrderDetails[i].Price
	}
	if result.Error != nil {
		log.Println(result.Error)
		return -1, result.Error
	}
	return grandtotal, nil
}
