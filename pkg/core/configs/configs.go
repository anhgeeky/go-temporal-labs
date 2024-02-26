package configs

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

func LoadConfig(path string) {
	viper.SetConfigFile(path)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	// viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func LoadConfigRoot(filename string) {
	LoadConfig(GetFileRootPath(filename))
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func RootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}

func GetFileRootPath(filePath string) string {
	fmt.Println("Root file path", filePath)
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		log.Fatalln("Does not exist file", filePath, err)
	}
	return path.Dir(filePath)
}
