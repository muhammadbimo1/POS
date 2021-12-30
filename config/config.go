package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Config struct {
	RouterEngine          *gin.Engine
	ApiBaseUrl            string
	RunMigration          string
	TableManagementConfig TableManagementConfig
	DataSourceName        string
}

type TableManagementConfig struct {
	ApiBaseUrl string
}

func NewConfig() *Config {
	config := new(Config)
	var dbHost = "localhost"
	RunMigration := "n"
	var dbPort = "5432"
	var dbName = "wmb_bimo_example"
	var dbUser = "postgres"
	var apiHost = "localhost"
	var apiport = "8080"
	var dbPassword = "12345678"
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	config.DataSourceName = dsn

	tableManagementBaseUrl := "http://localhost:8081/api/table"
	tableManagementConfig := TableManagementConfig{ApiBaseUrl: tableManagementBaseUrl}
	config.TableManagementConfig = tableManagementConfig

	r := gin.Default()
	config.RouterEngine = r
	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiport)
	config.RunMigration = RunMigration
	return config

}
