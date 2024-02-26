package activities

import (
	"context"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/moneytransfer"
	"go.temporal.io/sdk/activity"
)

type TransferActivity struct {
	AccountService       account.AccountService
	MoneyTransferService moneytransfer.MoneyTransferService
}

func (a *TransferActivity) CheckBalance(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CheckBalance", msg)

	res, err := a.AccountService.GetBalance()
	if err != nil {
		return err
	}

	logger.Info("TransferActivity: CheckBalance done", res)

	return nil
}

func (a *TransferActivity) CheckTargetAccount(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CheckTargetAccount", msg)
	return nil
}

func (a *TransferActivity) CreateTransferTransaction(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)

	// time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
	logger.Info("TransferActivity: CreateTransferTransaction", msg)
	res, err := a.MoneyTransferService.CreateTransferTransaction(msg.WorkflowID)
	if err != nil {
		return err
	}
	logger.Info("TransferActivity: CreateTransferTransaction done", res)

	return nil
}

func (a *TransferActivity) WriteCreditAccount(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	// time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
	logger.Info("TransferActivity: WriteCreditAccount", msg)
	res, err := a.MoneyTransferService.WriteCreditAccount(msg.WorkflowID)
	if err != nil {
		return err
	}
	logger.Info("TransferActivity: WriteCreditAccount done", res)
	return nil
}

func (a *TransferActivity) WriteDebitAccount(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	// time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
	logger.Info("TransferActivity: WriteDebitAccount", msg)
	res, err := a.MoneyTransferService.WriteDebitAccount(msg.WorkflowID)
	if err != nil {
		return err
	}
	logger.Info("TransferActivity: WriteDebitAccount done", res)
	return nil
}

func (a *TransferActivity) AddNewActivity(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	// time.Sleep(time.Duration(30) * time.Minute) // TODO: Test chờ 30p
	logger.Info("TransferActivity: AddNewActivity", msg)
	res, err := a.MoneyTransferService.AddNewActivity(msg.WorkflowID)
	if err != nil {
		return err
	}
	logger.Info("TransferActivity: AddNewActivity done", res)
	return nil
}

// ============================================
// Rollback
// ============================================

func (a *TransferActivity) CreateTransferTransactionCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: CreateTransferTransaction", msg)
	res, err := a.MoneyTransferService.CreateTransferTransaction(msg.WorkflowID)
	if err != nil {
		return err
	}
	logger.Info("TransferActivity: CreateTransferTransaction done", res)
	return nil
}

func (a *TransferActivity) WriteCreditAccountCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: WriteCreditAccountCompensation", msg)
	res, err := a.MoneyTransferService.WriteCreditAccountCompensation(msg.WorkflowID)
	if err != nil {
		return err
	}
	logger.Info("TransferActivity: WriteCreditAccountCompensation done", res)
	return nil
}

func (a *TransferActivity) WriteDebitAccountCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: WriteDebitAccountCompensation", msg)
	res, err := a.MoneyTransferService.WriteDebitAccountCompensation(msg.WorkflowID)
	if err != nil {
		return err
	}
	logger.Info("TransferActivity: WriteDebitAccountCompensation done", res)
	return nil
}

func (a *TransferActivity) AddNewActivityCompensation(ctx context.Context, msg messages.Transfer) error {
	logger := activity.GetLogger(ctx)
	logger.Info("TransferActivity: AddNewActivityCompensation", msg)
	res, err := a.MoneyTransferService.AddNewActivityCompensation(msg.WorkflowID)
	if err != nil {
		return err
	}
	logger.Info("TransferActivity: AddNewActivityCompensation done", res)
	return nil
}
