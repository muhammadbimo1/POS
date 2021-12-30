package usecase

import (
	"table_management/dto"
	"table_management/repository"
)

type TableReservationUseCase interface {
	ReserveTable(tableRequest dto.TableRequest) error
	CloseTable(billNo string) error
}

type tableReservationUseCase struct {
	reservationRepo repository.TableReservationRepository
}

func NewTableReservationUseCase(repo repository.TableReservationRepository) TableReservationUseCase {
	return &tableReservationUseCase{reservationRepo: repo}
}

func (t *tableReservationUseCase) ReserveTable(table dto.TableRequest) error {
	return t.reservationRepo.CallTableCheckIn(table)
}

func (t *tableReservationUseCase) CloseTable(billNo string) error {
	return t.reservationRepo.CallTableCheckOut(billNo)
}
