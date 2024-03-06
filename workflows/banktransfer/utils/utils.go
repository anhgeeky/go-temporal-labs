package utils

import (
	"fmt"
	"os"
)

func GetConsumerGroup(workflowId, activityId string) string {
	name, err := os.Hostname()

	// Có lỗi gắn default `::1`
	if err != nil {
		name = "::1"
	}

	// follow: "NEW-MCS-TEMPORAL-GO_WORKER_{HOSTNAME|POD NAME}_WORKFLOW_{WF_ID}_ACTIVITY_{ACT_ID}"
	return fmt.Sprintf("NEW-MCS-TEMPORAL-GO_WORKER_%s_WORKFLOW_%s_ACTIVITY_%s", name, workflowId, activityId)
}
