package activities

import (
	"context"

	"github.com/anhgeeky/go-temporal-labs/notification/messages"
	"github.com/anhgeeky/go-temporal-labs/notification/outbound/notification"
	"github.com/google/uuid"
	"go.temporal.io/sdk/activity"
)

type NotificationActivity struct {
	NotificationService notification.NotificationService
}

func (a *NotificationActivity) GetDeviceToken(ctx context.Context, msg interface{}) (*messages.DeviceToken, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: GetDeviceToken", msg)

	token := messages.DeviceToken{
		FirebaseToken: uuid.New().String(),
	}

	return &token, nil
}

func (a *NotificationActivity) PushEmail(ctx context.Context, msg *messages.DeviceToken) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: PushEmail", msg)

	return "OK", nil
}

func (a *NotificationActivity) PushSMS(ctx context.Context, msg *messages.DeviceToken) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: PushSMS", msg)

	return "OK", nil
}

func (a *NotificationActivity) PushNotification(ctx context.Context, msg *messages.DeviceToken) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: PushNotification", msg)

	return "OK", nil
}

func (a *NotificationActivity) PushInternalApp(ctx context.Context, msg *messages.DeviceToken) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: PushInternalApp", msg)

	return "OK", nil
}
