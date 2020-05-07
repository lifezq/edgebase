package service

import (
	"net/http"

	"code.uni-ledger.com/switch/edgebase/control/telemetry"
	"code.uni-ledger.com/switch/edgebase/internal/mqtt"
	"code.uni-ledger.com/switch/edgebase/internal/types"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/edgelist", getEdge)
	r.GET("/sendcmd", sendcmd)

	return r
}

func getEdge(c *gin.Context) {
	ed := telemetry.GetEdge()
	c.String(http.StatusOK, "%s", ed)
}

func sendcmd(c *gin.Context) {
	target := c.Query("target")
	dt := c.Query("cmd")
	msg := types.Cmd{
		ID:      mqtt.GetClientID(),
		MsgType: types.MSG_DOWNLOAD,
		Cmd:     []byte(dt),
	}

	ret, err := telemetry.SendMsg(target, &msg, 30)
	if err != nil {
		c.String(http.StatusOK, "%s", err.Error())
	}
	c.String(http.StatusOK, "%s", ret)
}
