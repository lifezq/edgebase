package telemetry

import (
	"encoding/json"
	"fmt"
	"time"

	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"code.uni-ledger.com/switch/edgebase/internal/types"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	edgeMap    map[string]string
	serviceMap map[string]string
)

func Init() {
	edgeMap = map[string]string{}
	token := mqtt.GetClient().Subscribe(types.TP_HeartBeat+"#", 0, handleHeartBeat)
	if token.Error() != nil {
		panic(token.Error())
	}

	token = mqtt.GetClient().Subscribe(types.TP_HeartBeatService+"#", 0, handleHeartBeatService)
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

func handleHeartBeatService(client MQTT.Client, msg MQTT.Message) {
	dt := types.HeartBeatService{}
	err := json.Unmarshal(msg.Payload(), &dt)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	serviceMap[dt.Service] = string(msg.Payload())
}

func GetEdge() []byte {

	var (
		em        = make(map[string]string)
		dt        = types.HeartBeat{}
		hbt int64 = 2
	)

	for client, payload := range edgeMap {

		err := json.Unmarshal([]byte(payload), &dt)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return []byte{}
		}

		if time.Now().Unix()-dt.Time < hbt+1 {
			em[client] = payload
		}
	}

	ret, _ := json.Marshal(em)
	return ret
}

func GetService() []byte {

	var (
		em        = make(map[string]string)
		dt        = types.HeartBeatService{}
		hbt int64 = 2
	)

	for client, payload := range serviceMap {

		err := json.Unmarshal([]byte(payload), &dt)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return []byte{}
		}

		if time.Now().Unix()-dt.Time < hbt+1 {
			em[client] = payload
		}
	}

	ret, _ := json.Marshal(em)
	return ret
}
