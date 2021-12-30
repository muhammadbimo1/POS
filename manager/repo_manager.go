package manager

import "table_management/repository"

type RepoManager interface {
	OrderRepo() repository.OrderRepository
	TableReservationRepo() repository.TableReservationRepository
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

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
