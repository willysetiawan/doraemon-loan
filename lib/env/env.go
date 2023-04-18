package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// // Bool defines a bool environment variable with specified name, and default value.
// // The return value is a bool variable that stores the value of the environment variable.

func init() {
	var fileConfig string
	fileConfig = "config.json"
	// switch os.Getenv("APP_ENV") {
	// case "DEV":
	// 	fileConfig = "env-dev.json"
	// case "STG":
	// 	fileConfig = "env-stg.json"
	// case "PROD":
	// 	fileConfig = "env-prod.json"
	// default:
	// 	fileConfig = "config.json"
	// }
	logrus.Info("APP ENV : " + os.Getenv("APP_ENV"))
	viper.SetConfigName(fileConfig)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warn(fmt.Sprintf("Error while reading config file %s", err))
	}
}

func String(key string, def string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		logrus.Warn(fmt.Sprintf("Error while reading config file %s", "Invalid type assertion - from = "+key))
		return def
	}
	return value
}

func Bool(key string, defaultValue bool) bool {
	if value, ok := viper.Get(key).(string); ok {
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			panic(fmt.Sprintf("%s value should be true or false, got %s", key, value))
		}
		return boolVal
	}

	return defaultValue
}

func Interface(key string, def interface{}) interface{} {
	value := viper.Get(key)
	return value
}
