package telemetry

import (
	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"code.uni-ledger.com/switch/edgebase/internal/types"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func Init() {
	token := mqtt.GetClient().Subscribe(types.TP_HeartBeat+"#", 0, handleHeartBeat)
	if token.Error() != nil {
		panic(token.Error())
	}
}

func handleHeartBeat(client MQTT.Client, msg MQTT.Message) {
	dt := types.HeartBeat{}
	err := json.Unmarshal(msg.Payload(), &dt)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	fmt.Println(msg.Topic(), ":", dt)
}
