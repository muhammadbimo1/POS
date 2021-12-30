package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerOrder struct {
	ID           string `gorm:"size:36;primaryKey"`
	CustomerName string `gorm:"size:50"`
	OrderDetails []CustomerOrderDetail
	gorm.Model
}

func (co *CustomerOrder) BeforeCreate(tx *gorm.DB) error {
	co.ID = uuid.NewString()
	return nil
}

func (co *CustomerOrder) TableName() string {
	return "customer_order"
}
