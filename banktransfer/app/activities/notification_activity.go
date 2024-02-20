package activities

import (
	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
	"go.temporal.io/sdk/workflow"
)

type NotificationActivity struct {
	// - 2.7.1 Lấy thông tin `token` của các thiết bị theo tài khoản
	// - 2.7.2 Push message SMS thông báo đã `Chuyển tiền Thành công`
	// - 2.7.3 Push message notification vào `firebase`
	// - 2.7.4 Push message internal app, reload lại màn hình hiện tại `Đang xử lý` -> `Thành công`
}

func (a *NotificationActivity) GetDeviceToken(ctx workflow.Context, msg interface{}) (*messages.DeviceToken, error) {
	logger := workflow.GetLogger(ctx)

	logger.Info("NotificationActivity: GetDeviceToken", msg)

	token := messages.DeviceToken{}

	return &token, nil
}

func (a *NotificationActivity) PushSMS(ctx workflow.Context, msg interface{}) error {
	logger := workflow.GetLogger(ctx)

	logger.Info("NotificationActivity: PushSMS", msg)

	return nil
}

func (a *NotificationActivity) PushNotification(ctx workflow.Context, msg interface{}) error {
	logger := workflow.GetLogger(ctx)

	logger.Info("NotificationActivity: PushNotification", msg)

	return nil
}

func (a *NotificationActivity) PushInternalApp(ctx workflow.Context, msg interface{}) error {
	logger := workflow.GetLogger(ctx)

	logger.Info("NotificationActivity: PushInternalApp", msg)

	return nil
}
