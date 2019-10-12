package internal

import (
	"github.com/spf13/viper"
)

func Configure() {
	viper.SetConfigName("config")
	viper.SetDefault("ConcurrentRequests", 1000)
	viper.SetDefault("TmpPath", "./tmp/")
	viper.SetDefault("FoundBooksPath", "./tmp/founded-books.json")
}
