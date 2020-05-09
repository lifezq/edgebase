package types

import (
	"encoding/json"
	"os"
)

const (
	MSG_DOWNLOAD = iota
	MSG_EXEC

	CMD_PATH = ".." + string(os.PathSeparator) + "cmd"
)

type Cmd struct {
	ID      string `json:"id"`
	MsgType int    `json:"msg_type`
	Cmd     []byte `json:cmd`
}

func GetCmd(c []byte) (*Cmd, error) {
	cmd := &Cmd{}
	err := json.Unmarshal(c, cmd)
	return cmd, err
}

const (
	FAIL    = 400
	SUCCESS = 200
)

type CmdRet struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewCmdRet() *CmdRet {
	return &CmdRet{
		Code: FAIL,
		Data: "not found",
	}
}
