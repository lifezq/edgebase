package telemetry

import (
	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"code.uni-ledger.com/switch/edgebase/internal/types"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func SendMsg(target string, msg []byte, wait int) ([]byte, error) {
	c := mqtt.GetClient()

	rtopic := types.TP_Return + mqtt.GetClientID()
	ch := make(chan []byte, 2)
	cb := func(client MQTT.Client, msg MQTT.Message) {
		ch <- msg.Payload()
	}

	t := c.Subscribe(rtopic, 2, cb)
	if t.Error() != nil {
		return nil, t.Error()
	}
	defer func() {
		if token := c.Unsubscribe(rtopic); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	}()

	t = c.Publish(types.TP_SendMsg+target, 2, false, msg)
	if t.Error() != nil {
		return nil, t.Error()
	}

	ret := []byte{}
	select {
	case <-time.After(time.Duration(wait) * time.Millisecond):
		return []byte{}, errors.New("time out")
	case dt := <-ch:
		ret = dt
	}

	return ret, nil
}
