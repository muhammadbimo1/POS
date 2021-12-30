package usecase

import (
	"errors"
	"log"
	"table_management/dto"
	"table_management/entity"
	"table_management/repository"
)

type OrderUseCase interface {
	OpenBill(orderRequest dto.OrderRequest) (*entity.CustomerOrder, error)
	CloseBill(closeBillRequest dto.CloseBillRequest) (string, error)
	GetTotal(billno string) (int, error)
}

type orderUseCase struct {
	orderRepo               repository.OrderRepository
	tableReservationUseCase TableReservationUseCase
	paymentUseCase          PaymentUseCase
}

func NewOrderUseCase(repo repository.OrderRepository, tableusecase TableReservationUseCase, paymentusecase PaymentUseCase) OrderUseCase {
	return &orderUseCase{
		orderRepo:               repo,
		tableReservationUseCase: tableusecase,
		paymentUseCase:          paymentusecase,
	}
}

func (o *orderUseCase) OpenBill(orderRequest dto.OrderRequest) (*entity.CustomerOrder, error) {
	order, err := o.orderRepo.CreateOne(*orderRequest.ToOrderEntity())
	if err != nil {
		return nil, errors.New("error on bill creation")
	}
	log.Println(order)
	tablerequest := dto.TableRequest{
		BillNo:  order.ID,
		TableId: orderRequest.TableId,
	}
	log.Println(tablerequest)
	err = o.tableReservationUseCase.ReserveTable(tablerequest)
	if err != nil {

		return order, errors.New("error on table request")
	}
	return order, err
}

func (o *orderUseCase) CloseBill(closeBillRequest dto.CloseBillRequest) (string, error) {
	bill, err := o.paymentUseCase.CreateBill(closeBillRequest)
	if err != nil {
		log.Println(err)
		return "", err
	}
	err = o.tableReservationUseCase.CloseTable(closeBillRequest.BillNo)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return bill.CustomerOrderId, nil
}

func (o *orderUseCase) GetTotal(billno string) (int, error) {
	return o.orderRepo.GetSum(billno)
}
