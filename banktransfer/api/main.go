package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
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

	app.Get("/accounts", controllers.GetAccountsHandler)
	app.Post("/transfers", CreateTransferHandler)
	app.Get("/transfers/:workflowID", GetTransferHandler)
	app.Put("/transfers/:workflowID/add", AddToTransferHandler)
	app.Put("/transfers/:workflowID/remove", RemoveFromTransferHandler)
	app.Put("/transfers/:workflowID/checkout", CheckoutHandler)
	app.Put("/transfers/:workflowID/email", UpdateEmailHandler)

	log.Println("App is running and listening on port", PORT)
	app.Listen(fmt.Sprintf(":%d", PORT))
}
