package main_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/anhgeeky/go-temporal-labs/core/configs"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_LoadConfig(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	filePath := filepath.Join(filepath.Dir(b), "../..", ".env")
	fmt.Println(filePath)

	configs.LoadConfig(filePath)

	assert.Equal(t, "3131", viper.GetString("PORT"))
	assert.Equal(t, "localhost:7233", viper.GetString("TEMPORAL_CLUSTER_HOST"))
	assert.Equal(t, "http://localhost:6001", viper.GetString("MCS_ACCOUNT_HOST"))
	assert.Equal(t, "http://localhost:6002", viper.GetString("MCS_MONEY_TRANSFER_HOST"))
	assert.Equal(t, "http://localhost:6003", viper.GetString("MCS_NOTIFICATION_HOST"))
}
