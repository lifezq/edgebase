package main

import (
	"code.uni-ledger.com/switch/edgebase/control/telemetry"
	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"fmt"
	"time"

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
	telemetry.Init()
}

func main() {
	time.Sleep(5 * time.Second)
	msg := []byte("cmd")
	ret, err := telemetry.SendMsg("edge1", msg, 5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("done:", ret)
	ch := make(chan struct{})
	<-ch
}
