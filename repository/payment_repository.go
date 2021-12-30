package repository

import (
	"log"
	"table_management/entity"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreateOne(order entity.OrderPayment) (*entity.OrderPayment, error)
	GetPaymentInfo(id string) (*entity.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func (p *paymentRepository) CreateOne(order entity.OrderPayment) (*entity.OrderPayment, error) {
	err := p.db.Create(&order).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &order, nil
}

func (p *paymentRepository) GetPaymentInfo(id string) (*entity.Payment, error) {
	var payment entity.Payment
	err := p.db.First(&payment, "id=?", id).Error
	if err != nil {
		log.Println(err)
		return &entity.Payment{}, err
	}
	return &payment, nil
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}
