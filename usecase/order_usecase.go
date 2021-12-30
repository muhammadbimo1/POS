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
}

type orderUseCase struct {
	orderRepo               repository.OrderRepository
	tableReservationUseCase TableReservationUseCase
}

func NewOrderUseCase(repo repository.OrderRepository, tableusecase TableReservationUseCase) OrderUseCase {
	return &orderUseCase{
		orderRepo:               repo,
		tableReservationUseCase: tableusecase,
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
	err := o.tableReservationUseCase.CloseTable(closeBillRequest.BillNo)
	if err != nil {
		return "", err
	}
	return closeBillRequest.BillNo, nil
}
