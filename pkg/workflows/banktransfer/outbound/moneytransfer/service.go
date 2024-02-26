package moneytransfer

import (
	"fmt"
	"log"

	"github.com/anhgeeky/go-temporal-labs/core/models"
	"github.com/go-resty/resty/v2"
)

var (
	ROUTE = "transfers"
)

type MoneyTransferService struct {
	Host string
	// Kafka
}

func (r MoneyTransferService) CreateTransferTransaction(workflowID string) (interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s/transfer", r.Host, ROUTE)
	client := resty.New()

	url := fmt.Sprintf("%s/%s/transactions", endpoint, workflowID)
	var data models.Response[string]

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&data).
		Post(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST Response:", response.Status())
	fmt.Printf("Retrieved %v \n", data)

	return response.Result(), err
}

func (r MoneyTransferService) WriteCreditAccount(workflowID string) (interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s/transfer", r.Host, ROUTE)
	client := resty.New()

	url := fmt.Sprintf("%s/%s/credit-accounts", endpoint, workflowID)
	var data models.Response[string]

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&data).
		Post(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST Response:", response.Status())
	fmt.Printf("Retrieved %v \n", data)

	return response.Result(), err
}

func (r MoneyTransferService) WriteDebitAccount(workflowID string) (interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s/transfer", r.Host, ROUTE)
	client := resty.New()

	url := fmt.Sprintf("%s/%s/debit-accounts", endpoint, workflowID)
	var data models.Response[string]

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&data).
		Post(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST Response:", response.Status())
	fmt.Printf("Retrieved %v \n", data)

	return response.Result(), err
}

func (r MoneyTransferService) AddNewActivity(workflowID string) (interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s/transfer", r.Host, ROUTE)
	client := resty.New()

	url := fmt.Sprintf("%s/%s/new-activity", endpoint, workflowID)
	var data models.Response[string]

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&data).
		Post(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST Response:", response.Status())
	fmt.Printf("Retrieved %v \n", data)

	return response.Result(), err
}

// ========================================
// Rollback
// ========================================

func (r MoneyTransferService) CreateTransferTransactionCompensation(workflowID string) (interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s/transfer", r.Host, ROUTE)
	client := resty.New()

	url := fmt.Sprintf("%s/%s/transactions/rollback", endpoint, workflowID)
	var data models.Response[string]

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&data).
		Post(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST Response:", response.Status())
	fmt.Printf("Retrieved %v \n", data)

	return response.Result(), err
}

func (r MoneyTransferService) WriteCreditAccountCompensation(workflowID string) (interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s/transfer", r.Host, ROUTE)
	client := resty.New()

	url := fmt.Sprintf("%s/%s/credit-accounts/rollback", endpoint, workflowID)
	var data models.Response[string]

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&data).
		Post(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST Response:", response.Status())
	fmt.Printf("Retrieved %v \n", data)

	return response.Result(), err
}

func (r MoneyTransferService) WriteDebitAccountCompensation(workflowID string) (interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s/transfer", r.Host, ROUTE)
	client := resty.New()

	url := fmt.Sprintf("%s/%s/debit-accounts/rollback", endpoint, workflowID)
	var data models.Response[string]

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&data).
		Post(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST Response:", response.Status())
	fmt.Printf("Retrieved %v \n", data)

	return response.Result(), err
}

func (r MoneyTransferService) AddNewActivityCompensation(workflowID string) (interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s/transfer", r.Host, ROUTE)
	client := resty.New()

	url := fmt.Sprintf("%s/%s/new-activity/rollback", endpoint, workflowID)
	var data models.Response[string]

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&data).
		Post(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST Response:", response.Status())
	fmt.Printf("Retrieved %v \n", data)

	return response.Result(), err
}
