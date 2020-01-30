package main

import (
	"log"
	"time"

	"github.com/nazwa/legend-api/pkg/legend"
	"github.com/spf13/viper"
)

func main() {
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Config file not loaded: %s", err)
	}

	config := legend.ConnectionConfig{
		UserAgent: viper.GetString("legend.useragent"),
		Username:  viper.GetString("legend.username"),
		Password:  viper.GetString("legend.password"),
		Guid:      viper.GetString("legend.guid"),
		Endpoint:  viper.GetString("legend.endpoint"),
		MaxRetries:  10,
		HttpTimeout: 120 * time.Second,
	}

	apiClient := legend.NewApiClient(config)

	locations, err := apiClient.GetAllLocations()
	if err != nil {
		panic(err)
	}

	for _, location := range locations {
		log.Println(location.Name)
	}
}
