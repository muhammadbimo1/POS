package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"table_management/config"
	"table_management/dto"
)

type OpoPaymentRepository interface {
	Payment(phoneNo string, total int) (string, error)
}

type opoPaymentRepository struct {
	httpClient       *http.Client
	opoPaymentConfig config.OpoPaymentConfig
}

func NewOpoRepository(httpClient *http.Client, config config.OpoPaymentConfig) OpoPaymentRepository {
	return &opoPaymentRepository{
		httpClient:       httpClient,
		opoPaymentConfig: config,
	}
}

func (o *opoPaymentRepository) Payment(phoneNo string, total int) (string, error) {
	postBody, _ := json.Marshal(dto.OpoRequest{
		CustomerPhoneNo: phoneNo,
		Total:           total,
	})
	requestBody := bytes.NewBuffer(postBody)
	request, err := http.NewRequest("POST", o.opoPaymentConfig.ApiBaseUrl, requestBody)
	if err != nil {
		log.Println("error opo reuest")
		return "", errors.New("error making opo request")
	}
	request.Header.Set("Opo-Client-Key", o.opoPaymentConfig.ClientSecretKey)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response, err := o.httpClient.Do(request)
	if err != nil {
		log.Println("error getting response")
		return "", errors.New("error doing getting response from opo")
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Println("error getting response")
		return "", errors.New("error")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var responseMessage dto.OpoPaymentResponse
	err = json.Unmarshal(body, &responseMessage)
	if err != nil {
		log.Println("error unmarshal")
		return "", err
	}
	if responseMessage.Message == "FAILED" {
		return "", errors.New(responseMessage.Message)
	}
	return responseMessage.Message, nil
}
