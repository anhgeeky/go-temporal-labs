package services_test

import (
	"fmt"
	"testing"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/services"
)

var (
	service services.AccountService
)

func init() {
	service.Host = "http://localhost:6001/"
}

func Test_GetBalance(t *testing.T) {
	fmt.Println(service.GetBalance())
}
