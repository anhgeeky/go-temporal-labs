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
	// - 2.7.1 Lấy thông tin `token` của các thiết bị theo tài khoản
	// - 2.7.2 Push message SMS thông báo đã `Chuyển tiền Thành công`
	// - 2.7.3 Push message notification vào `firebase`
	// - 2.7.4 Push message internal app, reload lại màn hình hiện tại `Đang xử lý` -> `Thành công`
}

func (a *NotificationActivity) GetDeviceToken(ctx context.Context, msg interface{}) (*messages.DeviceToken, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("NotificationActivity: GetDeviceToken", msg)

	token := messages.DeviceToken{
		FirebaseToken: uuid.New().String(),
	}

	return &token, nil
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
