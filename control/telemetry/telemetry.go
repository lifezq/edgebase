package telemetry

import (
	"encoding/json"
	"fmt"

	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"code.uni-ledger.com/switch/edgebase/internal/types"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var edgeMap map[string]string

func Init() {
	edgeMap = map[string]string{}
	token := mqtt.GetClient().Subscribe(types.TP_HeartBeat+"#", 0, handleHeartBeat)
	if token.Error() != nil {
		panic(token.Error())
	}
}

func handleHeartBeat(client MQTT.Client, msg MQTT.Message) {
	dt := types.HeartBeat{}
	err := json.Unmarshal(msg.Payload(), &dt)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	edgeMap[dt.Client] = string(msg.Payload())
}

func GetEdge() []byte {
	dt, _ := json.Marshal(edgeMap)
	return dt
}
