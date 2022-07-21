package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	InitConfig = initConfig
)

func initConfig(fileName string, additionalDirs []string) (err error) {
	viper.SetConfigName(fileName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	for _, dir := range additionalDirs {
		viper.AddConfigPath(dir)
	}
	// Read configuration file from disk
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to load config ", err)
		return
	}

	// Create configuration
	viper.ConfigFileUsed()
	viper.WatchConfig()
	log.Println("successfully read config file ")
	return
}

func getConfigString(key string) string {
	return viper.GetString(key)
}

// as of now not using the int config parameters so commented the below func

// func getConfigInt(key string) int {
// 	return viper.GetInt(key)
// }

func getConfigDuration(key string) time.Duration {
	return viper.GetDuration(key)
}
