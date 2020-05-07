package communication

import (
	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"code.uni-ledger.com/switch/edgebase/internal/types"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func handleCmd(c MQTT.Client, msg MQTT.Message) {
	fmt.Printf("get msg from:%s[%s]", msg.Topic(), msg.Payload())
	dt := msg.Payload()
	tk := mqtt.GetClient().Publish(types.TP_Return+mqtt.GetClientID(), 2, false, dt)
	if tk.Error() != nil {
		fmt.Println(tk.Error())
		return
	}
}
