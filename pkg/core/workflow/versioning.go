package workflow

import (
	"context"

	"go.temporal.io/sdk/client"
)

// FAQ: https://docs.temporal.io/dev-guide/go/versioning
func UpdateLatestWorkerBuildId(c client.Client, taskQueue, compatibleBuildID, latestBuildID string) {
	ctx := context.Background()
	c.UpdateWorkerBuildIdCompatibility(ctx, &client.UpdateWorkerBuildIdCompatibilityOptions{
		TaskQueue: taskQueue,
		Operation: &client.BuildIDOpAddNewCompatibleVersion{
			BuildID:                   latestBuildID,     // Version mới nhất hiện tại -> Apply mới vào
			ExistingCompatibleBuildID: compatibleBuildID, // Version trước đó -> Vẫn còn tương thích
		},
	})
}
