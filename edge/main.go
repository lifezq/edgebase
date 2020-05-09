package main

import (
	"fmt"

	"code.uni-ledger.com/switch/edgebase/edge/communication"
	"code.uni-ledger.com/switch/edgebase/edge/transfer"
	"code.uni-ledger.com/switch/edgebase/internal/mqtt"

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

var yamlExample1 = []byte(`
URL: http://39.100.139.81:8080/ping
Method: GET
Headers:
  Cookie: 55

`)

var yamlExample2 = []byte(`
URL: http://39.100.139.81:8080/tsping
Method: POST
Headers:
  Cookie: 55
  password: public
Body: "dfafdsfdsa"
  
`)

func main() {
	ret, err := transfer.RPC_http(yamlExample1)
	fmt.Println("do:", string(ret), err)
	ret, err = transfer.RPC_http(yamlExample2)
	fmt.Println("do:", string(ret), err)
	ch := make(chan struct{})
	<-ch
}
