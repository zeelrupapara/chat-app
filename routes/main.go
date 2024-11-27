package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeelrupapara/chat/config"
	"github.com/zeelrupapara/chat/websocket"
)

func Setup(router *gin.Engine, cfg config.AppConfig) {
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"host": cfg.Host,
			"port": cfg.Port,
		})
	})

	router.GET("/ws", func(c *gin.Context) {
		websocket.HandleConnections(c.Writer, c.Request)
	})
}
