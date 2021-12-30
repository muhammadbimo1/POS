package api

import (
	"database/sql"
	"log"
	"table_management/config"
	"table_management/delivery"
	"table_management/entity"
	"table_management/manager"
)

type Server interface {
	Run()
}

type server struct {
	config  *config.Config
	infra   manager.Infra
	usecase manager.UseCaseManager
}

func (s *server) Run() {
	if !(s.config.RunMigration == "Y" || s.config.RunMigration == "y") {
		db, err := s.infra.SqlDb().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Panicln(err)
			}
		}(db)
		s.InitRouter()
		s.config.RouterEngine.Run(s.config.ApiBaseUrl)
		if err != nil {
			log.Panicln(err)
		}
	} else {
		db := s.infra.SqlDb()
		err := db.AutoMigrate(&entity.CustomerOrder{}, entity.CustomerOrderDetail{}, entity.Payment{}, entity.OrderPayment{})
		db.Unscoped().Where("id like ?", "%%").Delete(entity.Payment{})
		db.Model(&entity.Payment{}).Save([]entity.Payment{
			{
				ID:                "P01",
				PaymentMethodName: "Tunai",
			},
			{
				ID:                "P02",
				PaymentMethodName: "OPO",
			},
		})
		if err != nil {
			log.Panicln(err)
		}
	}

}

func (s *server) InitRouter() {
	publicRoute := s.config.RouterEngine.Group("/api")
	delivery.NewOrderApi(publicRoute, s.usecase.OrderUseCase())
}

func NewApiServer() Server {
	appconfig := config.NewConfig()
	infra := manager.NewInfra(appconfig)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUseCaseManager(repo)

	return &server{
		config:  appconfig,
		infra:   infra,
		usecase: usecase,
	}
}
