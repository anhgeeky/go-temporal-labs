package activities_test

import (
	"fmt"
	"testing"

	"github.com/anhgeeky/go-temporal-labs/banktransfer/utils"
	"github.com/stretchr/testify/assert"
)

func Test_GetConsumerGroup(t *testing.T) {
	hostname := "AnhGeeky-PC"
	workflowID := "BANK_TRANSFER-1709525114"
	activityID := "check-balance"
	group := utils.GetConsumerGroup("BANK_TRANSFER-1709525114", "check-balance")

	fmt.Println("TRACE: ", group)

	assert.Equal(t, fmt.Sprintf("NEW-MCS-TEMPORAL-GO_WORKER_%s_WORKFLOW_%s_ACTIVITY_%s", hostname, workflowID, activityID), group)
}
