package main

import (
	"fmt"
	"time"

	"code.uni-ledger.com/switch/edgebase/control/service"
	"code.uni-ledger.com/switch/edgebase/control/telemetry"
	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"code.uni-ledger.com/switch/edgebase/internal/types"

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
	msg := types.Cmd{
		ID:      mqtt.GetClientID(),
		MsgType: types.MSG_DOWNLOAD,
		Cmd:     []byte("ok"),
	}
	ret, err := telemetry.SendMsg("edge1", &msg, 30)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("done:", string(ret))

	err = service.Init().Run()
	if err != nil {
		panic(err)
	}

}
