package activities

import (
	"context"

	"github.com/anhgeeky/go-temporal-labs/notification/messages"
	"go.temporal.io/sdk/activity"
)

type TransferActivity struct {
	// [**Transfer Flow**] Tạo lệnh YC chuyển tiền (`Create bank transfer`) (`Start`)
	// - 2.1. Kiểm tra số dư (`Check balance account`) (`Parallel`)
	// - 2.2. Kiểm tra tra tài khoản đích (`Check target account`) (`Parallel`)
	// - 2.3. Tạo giao dịch chuyển tiền (`Create new transaction`) (`When step 2.1, 2.2 done -> Continue`)
	// - 2.4. Tạo giao dịch ghi nợ (`Parallel`)
	// - 2.5. Tạo giao dịch ghi có (`Parallel`)
	// - 2.6. Transfer done  (`When step 2.4, 2.5 done -> Completed`) (`Trigger [Notification Flow]`)
	// - 2.7. Call subflow [**Notification Flow**] Gửi thông báo đã chuyển tiền
	// 	- 2.7.1 Lấy thông tin `token` của các thiết bị theo tài khoản
	// 	- 2.7.2 Push message SMS thông báo đã `Chuyển tiền Thành công`
	// 	- 2.7.3 Push message notification vào `firebase`
	// 	- 2.7.4 Push message internal app, reload lại màn hình hiện tại `Đang xử lý` -> `Thành công`
}

func (a *TransferActivity) CreateTransfer(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CreateTransfer", msg)
	return nil
}

func (a *TransferActivity) CheckBalance(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CheckBalance", msg)
	return nil
}

func (a *TransferActivity) CheckTargetAccount(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CheckTargetAccount", msg)
	return nil
}

func (a *TransferActivity) CreateTransferTransaction(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CreateTransferTransaction", msg)
	return nil
}

func (a *TransferActivity) WriteCreditAccount(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: WriteCreditAccount", msg)
	return nil
}

func (a *TransferActivity) WriteDebitAccount(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: WriteDebitAccount", msg)
	return nil
}

// ============================================
// Rollback
// ============================================
func (a *TransferActivity) CreateTransferTransactionCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CreateTransferTransaction", msg)
	return nil
}

func (a *TransferActivity) WriteCreditAccountCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: WriteCreditAccountCompensation", msg)
	return nil
}

func (a *TransferActivity) WriteDebitAccountCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: WriteDebitAccountCompensation", msg)
	return nil
}

// func StepWithError(ctx context.Context, transferDetails TransferDetails) error {
// 	fmt.Printf(
// 		"\nSimulate failure to trigger compensation. ReferenceId: %s\n",
// 		transferDetails.ReferenceID,
// 	)

// 	return errors.New("some error")
// }
