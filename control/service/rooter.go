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
	r.GET("/servicelist", getService)
	r.POST("/sendcmd", sendcmd)
	r.POST("/execcmd", execcmd)

	return r
}

func getEdge(c *gin.Context) {
	ed := telemetry.GetEdge()
	c.String(http.StatusOK, "%s", ed)
}

func getService(c *gin.Context) {
	ed := telemetry.GetService()
	c.String(http.StatusOK, "%s", ed)
}

func sendcmd(c *gin.Context) {

	target := c.PostForm("target")
	dt := c.PostForm("cmd")
	msg := types.Cmd{
		ID:      mqtt.GetClientID(),
		MsgType: types.MSG_DOWNLOAD,
		Cmd:     []byte(dt),
	}

	ret, err := telemetry.SendMsg(target, &msg, 60*1000*10)
	if err != nil {
		c.String(http.StatusOK, "%s", err.Error())
	}
	c.String(http.StatusOK, "%s", ret)
}

func execcmd(c *gin.Context) {

	target := c.PostForm("target")
	dt := c.PostForm("cmd")
	msg := types.Cmd{
		ID:      mqtt.GetClientID(),
		MsgType: types.MSG_EXEC,
		Cmd:     []byte(dt),
	}

	ret, err := telemetry.SendMsg(target, &msg, 30*1000)
	if err != nil {
		c.String(http.StatusOK, "%s", err.Error())
	}
	c.String(http.StatusOK, "%s", ret)
}
