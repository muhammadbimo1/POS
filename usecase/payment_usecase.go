package usecase

import (
	"errors"
	"log"
	"table_management/dto"
	"table_management/entity"
	"table_management/repository"
)

type PaymentUseCase interface {
	CreateBill(closeBillRequest dto.CloseBillRequest) (*entity.OrderPayment, error)
}

type paymentUseCase struct {
	paymentRepo repository.PaymentRepository
	opoRepo     repository.OpoPaymentRepository
}

func NewPaymentUseCase(repo repository.PaymentRepository, oporepo repository.OpoPaymentRepository) PaymentUseCase {
	return &paymentUseCase{repo, oporepo}
}

func (p *paymentUseCase) CreateBill(closeBillRequest dto.CloseBillRequest) (*entity.OrderPayment, error) {
	paymentInfo, err := p.paymentRepo.GetPaymentInfo(closeBillRequest.PaymentMethod)
	if err != nil {
		return nil, errors.New("error getting payment info")
	}
	var orderPaymentRequest = new(entity.OrderPayment)
	switch paymentInfo.ID {
	case "P02":
		opoPayment, err := p.opoPayment(closeBillRequest)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		orderPaymentRequest = opoPayment
	default:
		orderPaymentRequest = p.cashPayment(closeBillRequest)
	}
	return p.paymentRepo.CreateOne(*orderPaymentRequest)
}

func (p *paymentUseCase) cashPayment(payinfo dto.CloseBillRequest) *entity.OrderPayment {
	return &entity.OrderPayment{
		CustomerOrderId: payinfo.BillNo,
		BillerReceipt:   "",
		PaymentID:       payinfo.PaymentMethod,
	}
}

func (p *paymentUseCase) opoPayment(closeBillRequest dto.CloseBillRequest) (*entity.OrderPayment, error) {
	log.Print(closeBillRequest)
	receipt, err := p.opoRepo.Payment(closeBillRequest.PhoneNo, closeBillRequest.Total)
	if err != nil {
		return nil, errors.New("error on opoPayment")
	}
	return &entity.OrderPayment{
		CustomerOrderId: closeBillRequest.BillNo,
		BillerReceipt:   receipt,
		PaymentID:       closeBillRequest.PaymentMethod,
	}, nil
}
