package router

import (
	"chatgo/src/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	ApiRouter := gin.Default()
	ApiRouter.LoadHTMLGlob("frontend/html/*")
	ApiRouter.StaticFS("/frontend", http.Dir("./frontend"))

	ApiRouter.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", struct{}{})
	})

	ApiRouter.GET("/chat", func(c *gin.Context) {
		data := gin.H{
			"title": "chatgo",
		}
		c.HTML(http.StatusOK, "chat.html", data)
	})

	ApiRouter.GET("/ws", api.Ws)

	return ApiRouter
}
