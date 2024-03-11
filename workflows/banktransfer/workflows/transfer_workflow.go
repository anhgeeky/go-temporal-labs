package workflows

import (
	"fmt"
	"time"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/activities"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/config"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/messages"
	"github.com/anhgeeky/go-temporal-labs/banktransfer/outbound/account"
	cw "github.com/anhgeeky/go-temporal-labs/core/workflow"
	notiMsg "github.com/anhgeeky/go-temporal-labs/notification/messages"
	notiWorkflows "github.com/anhgeeky/go-temporal-labs/notification/workflows"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"go.temporal.io/sdk/workflow"
)

func TransferWorkflow(ctx workflow.Context, state messages.Transfer) (err error) {
	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 2 * time.Minute,
		HeartbeatTimeout:    10 * time.Second,
		// StartToCloseTimeout: 60 * time.Minute,
		// HeartbeatTimeout:    60 * time.Minute,
		RetryPolicy: cw.WorkflowConfigs.RetryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	verifyOtpChannel := workflow.GetSignalChannel(ctx, config.SignalChannels.VERIFY_OTP_CHANNEL)
	createTransactionChannel := workflow.GetSignalChannel(ctx, config.SignalChannels.CREATE_TRANSACTION_CHANNEL)
	completed := false

	var a *activities.TransferActivity

	for {

		// ====================== Activity: CheckBalance ======================
		var checkBalanceRes *account.CheckBalanceRes
		err = workflow.ExecuteActivity(ctx, a.CheckBalance, state).Get(ctx, &checkBalanceRes)
		if err != nil {
			return err
		}
		// ====================== Activity: CheckBalance ======================

		// ====================== Activity: CreateOTP ======================
		var createOTPRes *account.CreateOTPRes
		err = workflow.ExecuteActivity(ctx, a.CreateOTP, state).Get(ctx, &createOTPRes)
		if err != nil {
			return err
		}
		// ====================== Activity: CreateOTP ======================

		selector := workflow.NewSelector(ctx)

		// ====================== Signal: Verified OTP ======================
		selector.AddReceive(verifyOtpChannel, func(c workflow.ReceiveChannel, _ bool) {

			var signal interface{}
			c.Receive(ctx, &signal)

			var message messages.VerifiedOtpSignal
			err = mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			// ====================== Activity: CreateTransaction ======================
			// Có kết quả tạo OTP thành công + Xác thực OTP thành công
			var createTransactionRes *account.CreateTransactionRes
			err = workflow.ExecuteActivity(ctx, a.CreateTransaction, state).Get(ctx, &createTransactionRes)
			if err != nil {
				logger.Error("Error execute activity CreateTransaction: %v", err)
				return
			}
			// Compensation
			// defer func() {
			// 	if err != nil {
			// 		errCompensation := workflow.ExecuteActivity(ctx, a.CreateTransactionCompensation, state).Get(ctx, nil)
			// 		err = multierr.Append(err, errCompensation)
			// 	}
			// }()
			// // ====================== Activity: CreateTransaction ======================
		})
		// ====================== Signal: Verified OTP ======================

		// ====================== Signal: Trả về kết quả Tạo giao dịch thành công ======================
		selector.AddReceive(createTransactionChannel, func(c workflow.ReceiveChannel, _ bool) {

			var signal interface{}
			c.Receive(ctx, &signal)

			var message messages.CreateTransactionSignal
			err = mapstructure.Decode(signal, &message)
			if err != nil {
				logger.Error("Invalid signal type %v", err)
				return
			}

			// ====================== Subflow: NotificationWorkflow ======================
			// Tạo giao dịch thành công (Từ Signal) -> Gửi thông báo
			execution := workflow.GetInfo(ctx).WorkflowExecution
			childID := fmt.Sprintf("NOTIFICATION: %v", execution.RunID)
			cwo := workflow.ChildWorkflowOptions{
				WorkflowID: childID,
			}
			ctx = workflow.WithChildOptions(ctx, cwo)
			msgNotfication := notiMsg.NotificationMessage{
				// TODO: Bổ sung payload
				Token: notiMsg.DeviceToken{
					FirebaseToken: uuid.New().String(),
				},
			}

			var result string
			err = workflow.ExecuteChildWorkflow(ctx, notiWorkflows.NotificationWorkflow, msgNotfication).Get(ctx, &result)
			if err != nil {
				logger.Error("Parent execution received child execution failure.", "Error", err)
				return
			}
			logger.Info("Parent execution completed.", "Result", result)

			// ====================== Subflow: NotificationWorkflow ======================
			completed = true
		})
		// ====================== Signal: Trả về kết quả Tạo giao dịch thành công ======================

		selector.Select(ctx)

		// Xử lý transfer hoàn tất
		if completed {
			break
		}
	}

	logger.Info("Workflow completed.")
	return
}
