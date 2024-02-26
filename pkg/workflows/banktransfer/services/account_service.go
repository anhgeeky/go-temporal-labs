package services

import (
	"fmt"
	"log"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/services/responses"
	"github.com/anhgeeky/go-temporal-labs/core/models"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

var (
	SVC_HOST = fmt.Sprintf("%s/%s", config.MCS_ACCOUNT_HOST, "accounts")
)

type AccountService struct {
}

func (r AccountService) GetBalance() (interface{}, error) {
	accId, _ := uuid.Parse("54892431-0a67-4b66-91c7-255d2321b224") // TODO: Sample for test
	client := resty.New()

	url := fmt.Sprintf("%s/%s/balance", SVC_HOST, accId.String())
	fmt.Printf("URL: %s", url)

	var data models.Response[responses.BalanceRes]

	response, err := client.R().
		EnableTrace().
		SetHeader("Accept", "application/json").
		SetResult(&data).
		Get(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET Response:", response.Status())
	fmt.Printf("Retrieved %v \n", data)

	return response.Result(), err
}
