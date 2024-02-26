package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/anhgeeky/go-temporal-labs/core/configs"
	"github.com/anhgeeky/go-temporal-labs/mcs-notification/apis/routes"
	"github.com/anhgeeky/go-temporal-labs/mcs-notification/config"
	"github.com/anhgeeky/go-temporal-labs/mcs-notification/modules"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"

	"go.temporal.io/sdk/client"
)

var (
	temporal client.Client
	PORT     string
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(b), ".env")
	configs.LoadConfig(filePath)
	PORT := viper.GetInt32("PORT")
	log.Println("PORT", PORT)

	var err error
	temporal, err = client.NewLazyClient(client.Options{
		HostPort: config.TEMPORAL_CLUSTER_HOST,
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

	services := modules.SetupServices()

	routes.StartNotificationRoute(app, temporal, services)

	log.Println("App is running and listening on port", PORT)
	app.Listen(fmt.Sprintf(":%d", PORT))
}
