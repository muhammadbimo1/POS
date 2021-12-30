package manager

import "table_management/usecase"

type UseCaseManager interface {
	OrderUseCase() usecase.OrderUseCase
	TableReservationUseCase() usecase.TableReservationUseCase
	PaymentUseCase() usecase.PaymentUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (uc *useCaseManager) OrderUseCase() usecase.OrderUseCase {
	return usecase.NewOrderUseCase(uc.repo.OrderRepo(), uc.TableReservationUseCase(), uc.PaymentUseCase())
}

func (uc *useCaseManager) TableReservationUseCase() usecase.TableReservationUseCase {
	return usecase.NewTableReservationUseCase(uc.repo.TableReservationRepo())
}

func (uc *useCaseManager) PaymentUseCase() usecase.PaymentUseCase {
	return usecase.NewPaymentUseCase(uc.repo.PaymentRepo(), uc.repo.OpoRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: repoManager,
	}
}
