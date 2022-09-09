package api

import (
	"chatcord/src/ws"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Ws(c *gin.Context) {
	// Upgrade the http connection to webSocket Connection
	conn, err := ws.Upgrade(c.Writer, c.Request)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
	} else {
		zap.S().Infof("ws link at %s", conn.RemoteAddr().String())
	}

	client := ws.NewClient(conn.RemoteAddr().String(), conn)

	// Using Go Routines to Async Handle the read/write operations
	go client.Read()

	go client.Write()
}
