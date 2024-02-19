package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	app "github.com/anhgeeky/go-temporal-labs/bank-transfer"
	"github.com/anhgeeky/go-temporal-labs/bank-transfer/config"
	"github.com/anhgeeky/go-temporal-labs/bank-transfer/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"

	"go.temporal.io/sdk/client"
)

type (
	ErrorResponse struct {
		Message string
	}

	UpdateEmailRequest struct {
		Email string
	}

	CheckoutRequest struct {
		Email string
	}
)

var (
	temporal client.Client
	PORT     string
)

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	PORT := viper.GetInt32("PORT")
	log.Println("PORT", PORT)

	temporal, err = client.NewLazyClient(client.Options{
		HostPort: config.TemporalHost,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	log.Println("Temporal client connected")

	// middlewares
	app := fiber.New(fiber.Config{
		JSONDecoder: json.Unmarshal,
		JSONEncoder: json.Marshal,
	})

	// fiber log
	app.Use(logger.New(logger.Config{
		Next:         nil,
		Done:         nil,
		Format:       `${ip} - ${time} ${method} ${path} ${protocol} ${status} ${latency} "${ua}" "${error}"` + "\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stdout,
	}))

	app.Get("/accounts", GetAccountsHandler)
	app.Post("/bank-transfer", CreateTransferHandler)
	app.Get("/bank-transfer/:workflowID", GetTransferHandler)
	app.Put("/bank-transfer/:workflowID/add", AddToTransferHandler)
	app.Put("/bank-transfer/:workflowID/remove", RemoveFromTransferHandler)
	app.Put("/bank-transfer/:workflowID/checkout", CheckoutHandler)
	app.Put("/bank-transfer/:workflowID/email", UpdateEmailHandler)

	log.Println("App is running and listening on port", PORT)
	app.Listen(fmt.Sprintf(":%d", PORT))
}

func GetAccountsHandler(c *fiber.Ctx) error {
	res := domain.AccountList{}
	res.Accounts = domain.Accounts

	return c.Status(fiber.StatusOK).JSON(res)
}

func CreateTransferHandler(c *fiber.Ctx) error {
	workflowID := "TRANSFER-" + fmt.Sprintf("%d", time.Now().Unix())

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: app.Workflows.BANK_TRANSFER,
	}

	cart := app.TransferState{Items: make([]app.TransferItem, 0)}
	we, err := temporal.ExecuteWorkflow(context.Background(), options, app.TransferWorkflow, cart)
	if err != nil {
		return WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["cart"] = cart
	res["workflowID"] = we.GetID()

	return c.Status(fiber.StatusCreated).JSON(res)
}

func GetTransferHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	response, err := temporal.QueryWorkflow(context.Background(), workflowID, "", "getTransfer")
	if err != nil {
		return WriteError(c, err)
	}
	var res interface{}
	if err := response.Get(&res); err != nil {
		return WriteError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func AddToTransferHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	var item app.TransferItem
	json.Unmarshal(c.Body(), &item)

	update := app.AddToTransferSignal{Route: app.RouteTypes.ADD_TO_TRANSFER, Item: item}

	err := temporal.SignalWorkflow(context.Background(), workflowID, "", app.SignalChannels.ADD_TO_TRANSFER_CHANNEL, update)
	if err != nil {
		return WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}

func RemoveFromTransferHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")
	var item app.TransferItem
	json.Unmarshal(c.Body(), &item)

	update := app.RemoveFromTransferSignal{Route: app.RouteTypes.REMOVE_FROM_TRANSFER, Item: item}

	err := temporal.SignalWorkflow(context.Background(), workflowID, "", app.SignalChannels.REMOVE_FROM_TRANSFER_CHANNEL, update)
	if err != nil {
		return WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}

func UpdateEmailHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")

	var body UpdateEmailRequest
	json.Unmarshal(c.Body(), &body)
	updateEmail := app.UpdateEmailSignal{Route: app.RouteTypes.UPDATE_EMAIL, Email: body.Email}

	err := temporal.SignalWorkflow(context.Background(), workflowID, "", app.SignalChannels.UPDATE_EMAIL_CHANNEL, updateEmail)
	if err != nil {
		return WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["ok"] = 1
	return c.Status(fiber.StatusOK).JSON(res)
}

func CheckoutHandler(c *fiber.Ctx) error {
	workflowID := c.Params("workflowID")

	var body CheckoutRequest
	json.Unmarshal(c.Body(), &body)
	checkout := app.CheckoutSignal{Route: app.RouteTypes.CHECKOUT, Email: body.Email}

	err := temporal.SignalWorkflow(context.Background(), workflowID, "", app.SignalChannels.CHECKOUT_CHANNEL, checkout)
	if err != nil {
		return WriteError(c, err)
	}

	res := make(map[string]interface{})
	res["sent"] = true
	return c.Status(fiber.StatusOK).JSON(res)
}

func NotFoundHandler(c *fiber.Ctx) error {
	res := ErrorResponse{Message: "Endpoint not found"}
	return c.Status(fiber.StatusNotFound).JSON(res)
}

func WriteError(c *fiber.Ctx, err error) error {
	res := ErrorResponse{Message: err.Error()}
	return c.Status(fiber.StatusInternalServerError).JSON(res)
}
