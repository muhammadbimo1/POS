package manager

import "table_management/repository"

type RepoManager interface {
	OrderRepo() repository.OrderRepository
	TableReservationRepo() repository.TableReservationRepository
	PaymentRepo() repository.PaymentRepository
	OpoRepo() repository.OpoPaymentRepository
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) OrderRepo() repository.OrderRepository {
	return repository.NewOrderRepository(r.infra.SqlDb())
}

func (r *repoManager) TableReservationRepo() repository.TableReservationRepository {
	return repository.NewTableRepository(r.infra.HttpClient(), r.infra.Config().TableManagementConfig)
}

func (r *repoManager) PaymentRepo() repository.PaymentRepository {
	return repository.NewPaymentRepository(r.infra.SqlDb())
}

func (r *repoManager) OpoRepo() repository.OpoPaymentRepository {
	return repository.NewOpoRepository(r.infra.HttpClient(), r.infra.Config().OpoPaymentConfig)
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
