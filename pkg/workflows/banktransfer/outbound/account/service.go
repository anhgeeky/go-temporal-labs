package account

import (
	"fmt"
	"log"

	"github.com/anhgeeky/go-temporal-labs/core/models"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

var (
	ROUTE = "accounts"
)

type AccountService struct {
	Host string
}

func (r AccountService) GetBalance() (interface{}, error) {
	endpoint := fmt.Sprintf("%s/%s", r.Host, ROUTE)
	accId, _ := uuid.Parse("54892431-0a67-4b66-91c7-255d2321b224") // TODO: Sample for test
	client := resty.New()

	url := fmt.Sprintf("%s/%s/balance", endpoint, accId.String())
	var data models.Response[CheckBalanceRes]

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
