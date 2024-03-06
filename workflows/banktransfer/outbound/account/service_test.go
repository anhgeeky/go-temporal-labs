package account_test

import (
	"fmt"
	"testing"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
)

var (
	service account.AccountService
)

func init() {
	service.Host = "http://localhost:6001/"
}

func Test_GetBalance(t *testing.T) {
	fmt.Println(service.GetBalance())
}
