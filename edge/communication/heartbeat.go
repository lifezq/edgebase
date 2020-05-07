package communication

import (
	"encoding/json"
	"fmt"
	"time"

	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"code.uni-ledger.com/switch/edgebase/internal/types"
	"github.com/spf13/viper"
)

var state types.HeartBeat

func Init() {
	state = types.HeartBeat{
		Client: mqtt.GetClientID(),
		Alive:  true,
	}

	tk := mqtt.GetClient().Subscribe(types.TP_SendMsg+mqtt.GetClientID(), 2, handleCmd)
	if tk.Error() != nil {
		panic(tk.Error())
	}

	go sendHeartBeat()
}

func heartBeat(dt []byte) {
	token := mqtt.GetClient().Publish(types.TP_HeartBeat+mqtt.GetClientID(), 0, false, dt)
	if token.Error() != nil {
		fmt.Println(token.Error())
	}
}

func sendHeartBeat() {
	fmt.Println("Start communication")
	defer fmt.Println("End communication...")
	hbt := viper.GetInt64("broker.heartbeat")
	fmt.Println(hbt)
	tc := time.Tick(time.Duration(hbt) * time.Second)
	for {
		select {
		case <-tc:
			st := state
			st.Time = time.Now().Unix()
			dt, err := json.Marshal(st)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
			}
			heartBeat(dt)
		}
	}
}
