package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"table_management/config"
	"table_management/delivery/appresponse"
	"table_management/dto"
)

type TableReservationRepository interface {
	CallTableCheckIn(table dto.TableRequest) error
	CallTableCheckOut(billNo string) error
}

type tableReservationRepository struct {
	httpClient            *http.Client
	tableManagementConfig config.TableManagementConfig
}

func NewTableRepository(httpClient *http.Client, config config.TableManagementConfig) TableReservationRepository {
	return &tableReservationRepository{
		httpClient:            httpClient,
		tableManagementConfig: config,
	}
}

func (t *tableReservationRepository) CallTableCheckIn(table dto.TableRequest) error {
	log.Println("calltablecheckihn")
	postBody, _ := json.Marshal(table)
	log.Println(postBody)
	requestBody := bytes.NewBuffer(postBody)
	endPoint := fmt.Sprintf("%s/checkin", t.tableManagementConfig.ApiBaseUrl)
	log.Println(endPoint)
	response, err := t.httpClient.Post(endPoint, "application/json", requestBody)
	if err != nil {
		return errors.New("check in error")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("check in error")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var responseMessage appresponse.ResponseMessage
	err = json.Unmarshal(body, &responseMessage)
	if err != nil {
		return err
	}
	if responseMessage.Status == "FAILED" {
		return errors.New(responseMessage.Description)
	}
	return nil
}

func (t *tableReservationRepository) CallTableCheckOut(billNo string) error {
	endPoint := fmt.Sprintf("%s/checkout?billNo=%s", t.tableManagementConfig.ApiBaseUrl, billNo)
	request, err := http.NewRequest(http.MethodPut, endPoint, nil)
	if err != nil {
		return errors.New("check out error")
	}
	response, err := t.httpClient.Do(request)
	if err != nil {
		return errors.New("check out error")
	}
	if response.StatusCode == http.StatusOK {
		return nil
	}
	return errors.New("check out error")
}
