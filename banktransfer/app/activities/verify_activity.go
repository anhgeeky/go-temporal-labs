package activities

// import (
// 	"math/rand"
// 	"time"

// 	"github.com/anhgeeky/go-temporal-labs/banktransfer/app/messages"
// 	"go.temporal.io/sdk/workflow"
// )

// type VerifyActivity struct {
// }

// func (a *VerifyActivity) VerifyOtp(ctx workflow.Context, msg messages.VerifyOtpMessage) error {
// 	logger := workflow.GetLogger(ctx)

// 	logger.Info("VerifyActivity processing started.")

// 	// ================= SAMPLE CODE =================
// 	// Process: Verify `token`, `code`
// 	timeNeededToProcess := time.Second * time.Duration(rand.Intn(10))
// 	time.Sleep(timeNeededToProcess)
// 	// ================= SAMPLE CODE =================

// 	logger.Info("VerifyActivity done.", "duration", timeNeededToProcess)

// 	return nil
// }
