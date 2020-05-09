package communication

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
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
	case types.MSG_EXEC:
		result, err = handleExec(cmd.Cmd)
	}

	if err != nil {
		ret.Code = types.FAIL
		ret.Data = err.Error()
	} else {
		ret.Code = types.SUCCESS
		ret.Data = string(result)
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
	url := string(c)
	fmt.Println("get download cmd:", url)

	_, err := os.Stat(types.CMD_PATH)
	if err != nil && os.IsNotExist(err) {
		err := os.Mkdir(types.CMD_PATH, 0755)
		if err != nil {
			fmt.Printf("os.Mkdir error:%s\n", err.Error())
			return []byte(""), err
		}
	}

	fmt.Printf("http.get url:%s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	base := path.Base(url)
	filename := types.CMD_PATH + "/" + base
	fp, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0755)
	defer fp.Close()
	if err != nil {
		return []byte(""), err
	}

	fp.Write(body)

	if strings.Contains(base, ".tar.gz") {
		err := exec.Command("tar", "-zxvf", filename, "-C", types.CMD_PATH).Run()
		if err != nil {
			return []byte(""), err
		}
	} else if strings.Contains(base, ".tar") {
		err := exec.Command("tar", "-xvf", filename, "-C", types.CMD_PATH).Run()
		if err != nil {
			return []byte(""), err
		}
	} else if strings.Contains(base, ".zip") {
		err := exec.Command("unzip", filename, "-d", types.CMD_PATH).Run()
		if err != nil {
			return []byte(""), err
		}
	}

	err = exec.Command("rm", "-rf", filename).Run()
	if err != nil {
		return []byte(""), err
	}

	return []byte("ok"), nil
}

func handleExec(c []byte) ([]byte, error) {
	cmd := string(c)
	fmt.Println("get exec cmd:", cmd)

	ns := strings.Split(cmd, " ")

	_, err := os.Stat(ns[0])
	if err != nil && os.IsNotExist(err) {
		return exec.Command(ns[0], ns[1:]...).Output()
	}

	var (
		cancel  chan bool
		command *exec.Cmd
	)

	if strings.HasSuffix(ns[0], "sh") {
		command = exec.Command("sh", ns...)
	} else {
		command = exec.Command(ns[0], ns[1:]...)
	}

	defer func(service string, command *exec.Cmd, cancel chan bool) {
		err := command.Wait()
		if err != nil {
			fmt.Printf("service [%s] exit, err:%v\n", service, err)
			return
		}

		fmt.Printf("service [%s] exit, code:0\n", service)

		cancel <- true
	}(ns[0], command, cancel)
	err = command.Start()
	sendHeartBeatService(ns[0], cancel)
	if err != nil {
		return []byte(""), err
	}

	return []byte("ok"), nil
}
