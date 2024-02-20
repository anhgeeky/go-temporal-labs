package controllers

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/modules/account"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

type AccountController struct {
	Service        account.Service
	TemporalClient client.Client
}

func (r AccountController) GetAccounts(c *fiber.Ctx) error {
	res := account.AccountList{}

	items, err := r.Service.GetAccounts()
	if err != nil {
		return err
	}

	res.Accounts = *items

	return c.Status(fiber.StatusOK).JSON(res)
}

// TODO: Code
func (r AccountController) GetBalance(c *fiber.Ctx) error {
	res := account.AccountList{}

	items, err := r.Service.GetAccounts()
	if err != nil {
		return err
	}

	res.Accounts = *items

	return c.Status(fiber.StatusOK).JSON(res)
}
