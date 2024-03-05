package activities_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetConsumerGroup(t *testing.T) {
	hostname := "AnhGeeky-PC"
	workflowID := "BANK_TRANSFER-1709525114"
	activityID := "check-balance"
	group := getConsumerGroup("BANK_TRANSFER-1709525114", "check-balance")

	fmt.Println("TRACE: ", group)

	assert.Equal(t, fmt.Sprintf("NEW-MCS-TEMPORAL-GO_WORKER_%s_WORKFLOW_%s_ACTIVITY_%s", hostname, workflowID, activityID), group)
}

func getConsumerGroup(workflowId string, activityId string) string {
	name, err := os.Hostname()

	// Có lỗi gắn default `::1`
	if err != nil {
		name = "::1"
	}

	// follow: "NEW-MCS-TEMPORAL-GO_WORKER_{HOSTNAME|POD NAME}_WORKFLOW_{WF_ID}_ACTIVITY_{ACT_ID}"
	return fmt.Sprintf("NEW-MCS-TEMPORAL-GO_WORKER_%s_WORKFLOW_%s_ACTIVITY_%s", name, workflowId, activityId)
}
