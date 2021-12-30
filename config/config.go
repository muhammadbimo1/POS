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
	OpoPaymentConfig      OpoPaymentConfig
	DataSourceName        string
}

type TableManagementConfig struct {
	ApiBaseUrl string
}

type OpoPaymentConfig struct {
	ApiBaseUrl      string
	ClientSecretKey string
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
	opoBaseURL := "http://159.223.42.164:8899/opo/payment"
	secretKey := "E157934D-EA2E-49F6-9DCE-398B750BE4F0"
	tableManagementConfig := TableManagementConfig{ApiBaseUrl: tableManagementBaseUrl}
	opoPaymentConfig := OpoPaymentConfig{ApiBaseUrl: opoBaseURL, ClientSecretKey: secretKey}
	config.OpoPaymentConfig = opoPaymentConfig
	config.TableManagementConfig = tableManagementConfig
	r := gin.Default()
	config.RouterEngine = r
	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiport)
	config.RunMigration = RunMigration
	return config

}
