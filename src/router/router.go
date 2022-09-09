package router

import (
	"chatcord/src/api"
	"net/http"

	"github.com/gin-gonic/gin"
)
// Here we define the Routing Logic

func Routers() *gin.Engine {
	ApiRouter := gin.Default()
	ApiRouter.LoadHTMLGlob("frontend/html/*")
	ApiRouter.StaticFS("/frontend", http.Dir("./frontend"))

	// "/" will Route to Home Page
	ApiRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", struct{}{})
	})

	// Route to main Chat Page
	ApiRouter.GET("/chat", func(c *gin.Context) {
		data := gin.H{
			"title": "chatcord",
		}
		c.HTML(http.StatusOK, "chat.html", data)
	})

	// Reroute to Homepage
	ApiRouter.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/")
	})

	// Handling WebSocket Communication
	ApiRouter.GET("/ws", api.Ws)

	return ApiRouter
}
