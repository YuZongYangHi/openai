package routers

import (
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/controllers/channel"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/controllers/message"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	g := r.Group("/api/v1/openai")

	{
		g.GET("/channel/:channelId", channel.RetrieveController)
		g.DELETE("/channel/:channelId", channel.DeleteController)
		g.PUT("/channel/:channelId", channel.UpdateController)
		g.GET("/channel", channel.ListController)
		g.POST("/channel", channel.CreateController)

		g.GET("/message/:channelId", message.ListMessage)
		g.POST("/message/:channelId", message.CreateMessage)
	}
	return r
}
