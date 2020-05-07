package main

import (
	"code.uni-ledger.com/switch/edgebase/edge/communication"
	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	mqtt.Init()
	communication.Init()
}

func main() {

	ch := make(chan struct{})
	<-ch
}
