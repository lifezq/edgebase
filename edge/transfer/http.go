package transfer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func RPC_http(cmd []byte) ([]byte, error) {
	ret := []byte{}
	vp := viper.New()
	vp.SetConfigType("yaml")

	err := vp.ReadConfig(bytes.NewBuffer(cmd))
	if err != nil {
		panic(err)
	}

	bytesData := vp.GetString("Body")
	method := vp.GetString("Method")
	url := vp.GetString("URL")
	fmt.Println(bytesData, method, url)
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(bytesData)))
	if err != nil {
		return ret, err
	}

	hd := vp.GetStringMapString("Headers")
	for k, v := range hd {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}

	return ioutil.ReadAll(resp.Body)
}
