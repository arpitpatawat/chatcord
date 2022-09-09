package main

import (
	"chatcord/initial"
	"chatcord/src/router"
	"chatcord/src/ws"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	initial.InitLogger()

	Router := router.Routers()

	go ws.Manager.Start()

	zap.S().Infof("start server at: %d", 8000)

	if err := Router.Run(fmt.Sprintf(":%d", 8000)); err != nil {
		zap.S().Panic("start server failed: ", err.Error())
	}
}
