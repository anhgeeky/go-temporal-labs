package activities

import (
	"context"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"go.temporal.io/sdk/activity"
)

type TransferActivity struct{}

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
	time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
	return nil
}

func (a *TransferActivity) CreateTransferTransaction(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
	logger.Info("TransferActivity: CreateTransferTransaction", msg)
	return nil
}

func (a *TransferActivity) WriteCreditAccount(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
	logger.Info("TransferActivity: WriteCreditAccount", msg)
	return nil
}

func (a *TransferActivity) WriteDebitAccount(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
	logger.Info("TransferActivity: WriteDebitAccount", msg)
	return nil
}

func (a *TransferActivity) AddNewActivity(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
	logger.Info("TransferActivity: AddNewActivity", msg)
	return nil
}

// ============================================
// Rollback
// ============================================

func (a *TransferActivity) CheckBalanceCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CheckBalanceCompensation", msg)
	return nil
}

func (a *TransferActivity) CheckTargetAccountCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CheckTargetAccountCompensation", msg)
	return nil
}

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

func (a *TransferActivity) AddNewActivityCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: AddNewActivityCompensation", msg)
	return nil
}

// func StepWithError(ctx context.Context, transferDetails TransferDetails) error {
// 	fmt.Printf(
// 		"\nSimulate failure to trigger compensation. ReferenceId: %s\n",
// 		transferDetails.ReferenceID,
// 	)

// 	return errors.New("some error")
// }
