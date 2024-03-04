package controllers

import (
	"encoding/json"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/core/apis/responses"
	"github.com/anhgeeky/go-temporal-labs/core/broker"
	"github.com/anhgeeky/go-temporal-labs/mcs-account/modules/account"
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

	return responses.SuccessResult[account.AccountList](c, res)
}

func (r AccountController) GetBalance(c *fiber.Ctx) error {
	// res, err := r.Service.GetBalance()
	// if err != nil {
	// 	return err
	// }

	var req account.CheckBalanceReq
	json.Unmarshal(c.Body(), &req)

	res := account.BalanceRes{Balance: 9999}
	body, err := json.Marshal(res)
	if err != nil {
		return err
	}
	replyTopic := config.Messages.CHECK_BALANCE_REPLY_TOPIC

	fMsg := broker.Message{
		Body: body,
		Headers: map[string]string{
			"workflow_id": req.WorkflowID,
			"activity-id": req.Action,
		},
	}

	r.Service.Broker.Publish(replyTopic, &fMsg)

	return responses.SuccessResult[account.BalanceRes](c, res)
}
