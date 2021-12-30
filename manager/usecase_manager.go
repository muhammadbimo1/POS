package manager

import "table_management/usecase"

type UseCaseManager interface {
	OrderUseCase() usecase.OrderUseCase
	TableReservationUseCase() usecase.TableReservationUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (uc *useCaseManager) OrderUseCase() usecase.OrderUseCase {
	return usecase.NewOrderUseCase(uc.repo.OrderRepo(), uc.TableReservationUseCase())
}

func (uc *useCaseManager) TableReservationUseCase() usecase.TableReservationUseCase {
	return usecase.NewTableReservationUseCase(uc.repo.TableReservationRepo())
}

func NewUseCaseManager(repoManager RepoManager) UseCaseManager {
	return &useCaseManager{
		repo: repoManager,
	}
}
