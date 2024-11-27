package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zeelrupapara/chat/config"
	"github.com/zeelrupapara/chat/redis"
	"github.com/zeelrupapara/chat/routes"
	"github.com/zeelrupapara/chat/websocket"
)
func main() {
	cfg := config.GetConfig()
	fmt.Println(cfg)
	redis.InitRedis(cfg)

	router := gin.Default()

	routes.Setup(router, cfg)

	go websocket.ManageRoom()
	add := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	fmt.Printf("Listening on %s\n", add)
	if err := router.Run(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)); err != nil {
		panic(err)
	}
}
