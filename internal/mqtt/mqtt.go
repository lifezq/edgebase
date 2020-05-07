package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

var client MQTT.Client
var clientId string

func Init() {
	clientId = viper.GetString("broker.client")
	opts := MQTT.NewClientOptions().AddBroker(viper.GetString("broker.addr"))
	opts.SetClientID(clientId)
	opts.SetUsername(viper.GetString("broker.user"))
	opts.SetPassword(viper.GetString("broker.password"))
	//opts.SetDefaultPublishHandler(f)
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	client = c
}

func GetClient() MQTT.Client {
	return client
}

func GetClientID() string {
	return clientId
}
