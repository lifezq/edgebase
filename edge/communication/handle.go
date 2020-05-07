package communication

import (
	"encoding/json"
	"fmt"
	"strings"

	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"code.uni-ledger.com/switch/edgebase/internal/types"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func handleCmd(c MQTT.Client, msg MQTT.Message) {
	fmt.Printf("get msg from:%s[%s]\n", msg.Topic(), msg.Payload())

	ret := types.NewCmdRet()
	defer func() {
		retBy, _ := json.Marshal(ret)
		tp := types.TP_Return + strings.TrimPrefix(msg.Topic(), types.TP_SendMsg)
		fmt.Println("return topic:", tp)
		tk := mqtt.GetClient().Publish(tp, 2, false, retBy)
		if tk.Error() != nil {
			fmt.Println(tk.Error())
			return
		}
	}()

	dtBy := msg.Payload()
	cmd, err := types.GetCmd(dtBy)
	if err != nil {
		fmt.Println(err)
		return
	}

	result := []byte{}
	switch cmd.MsgType {
	case types.MSG_DOWNLOAD:
		result, err = handleDownload(cmd.Cmd)
	}

	if err != nil {
		ret.Code = types.FAIL
		ret.Data = err.Error()
	} else {
		ret.Code = types.SUCCESS
		ret.Data = result
	}

	retBy, _ := json.Marshal(ret)
	tp := types.TP_Return + cmd.ID
	fmt.Println("return topic:", tp)
	tk := mqtt.GetClient().Publish(tp, 2, false, retBy)
	if tk.Error() != nil {
		fmt.Println(tk.Error())
		return
	}
}

func handleDownload(c []byte) ([]byte, error) {
	fmt.Println("get download cmd:", string(c))
	return []byte("ok"), nil
}
