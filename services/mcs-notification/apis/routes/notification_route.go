package routes

import (
	"github.com/anhgeeky/go-temporal-labs/mcs-notification/apis/controllers"
	"github.com/anhgeeky/go-temporal-labs/mcs-notification/modules/notification"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func StartNotificationRoute(app *fiber.App, temporal client.Client, services map[string]interface{}) {
	controller := controllers.NotificationController{
		Service:        services["notificationService"].(notification.Service),
		TemporalClient: temporal,
	}
	group := app.Group("/notifications")
	group.Post("/", controller.NotificationOtp)
}
