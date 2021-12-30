package entity

import "gorm.io/gorm"

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
