package entity

import "gorm.io/gorm"

type Payment struct {
	ID                string `gorm:"column:id;size:3;primaryKey"`
	PaymentMethodName string `gorm:"size:20;"`
	gorm.Model
}

func (co *Payment) TableName() string {
	return "payment"
}
