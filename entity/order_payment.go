package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderPayment struct {
	CustomerOrderId string
	BillerReceipt   string `gorm:"size:50"`
	PaymentID       string
	Payment         Payment
	gorm.Model
}

func (co *OrderPayment) TableName() string {
	return "order_payment"
}

func (co *OrderPayment) BeforeCreate(tx *gorm.DB) error {
	co.ID = uint(uuid.New().ID())
	return nil
}
