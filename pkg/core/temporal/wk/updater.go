package wk

import (
	"context"
	"log"
	"sync"
	"time"

	"go.temporal.io/sdk/client"
)

// Thay đổi build worker running thành 1 version
// FAQ: https://docs.temporal.io/dev-guide/go/versioning
func UpdateWorkerBuildID(c client.Client, wg *sync.WaitGroup, taskQueue, buildID string) error {
	ctx := context.Background()
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := c.UpdateWorkerBuildIdCompatibility(ctx, &client.UpdateWorkerBuildIdCompatibilityOptions{
			TaskQueue: taskQueue,
			Operation: &client.BuildIDOpAddNewIDInNewDefaultSet{
				BuildID: buildID,
			},
		})
		if err != nil {
			log.Fatalf("Update worker: BuildID: %s, Trace: %v", buildID, err)
		}
		time.Sleep(5 * time.Second)

	}()

	return nil
}

// Thay đổi build worker running từ 1 version đến 1 version mới. Vẫn còn tương thích version trước đó
// FAQ: https://docs.temporal.io/dev-guide/go/versioning
func UpdateLatestWorkerBuildIDs(c client.Client, wg *sync.WaitGroup, taskQueue, compatibleBuildID, latestBuildID string) error {
	ctx := context.Background()
	wg.Add(1)
	go func() {
		defer wg.Done()

		var err error
		if compatibleBuildID == latestBuildID {
			err := c.UpdateWorkerBuildIdCompatibility(ctx, &client.UpdateWorkerBuildIdCompatibilityOptions{
				TaskQueue: taskQueue,
				Operation: &client.BuildIDOpAddNewIDInNewDefaultSet{
					BuildID: compatibleBuildID,
				},
			})
			if err != nil {
				log.Fatalf("Update worker: BuildID: %s, Trace: %v", compatibleBuildID, err)
			}
		} else {
			err = c.UpdateWorkerBuildIdCompatibility(ctx, &client.UpdateWorkerBuildIdCompatibilityOptions{
				TaskQueue: taskQueue,
				Operation: &client.BuildIDOpAddNewCompatibleVersion{
					BuildID:                   latestBuildID,     // Version mới nhất hiện tại -> Apply mới vào
					ExistingCompatibleBuildID: compatibleBuildID, // Version trước đó -> Vẫn còn tương thích
				},
			})
			if err != nil {
				log.Fatalf("Update worker: CompatibleBuildID: %s, LatestBuildID: %s. Trace: %v", compatibleBuildID, latestBuildID, err)
			}
		}

		time.Sleep(5 * time.Second)

	}()

	return nil
}
