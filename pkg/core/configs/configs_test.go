package configs_test

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func Test_LoadConfig(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(path.Dir(filename))
}
