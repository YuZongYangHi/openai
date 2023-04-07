package channel

import (
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/controllers/base"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/models"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/pkg/util/parsers"
	"github.com/gin-gonic/gin"
)

func ListController(ctx *gin.Context) {
	result := models.ChannelModel.List()
	base.SuccessResponse(ctx, result)
}

func CreateController(ctx *gin.Context) {
	var form models.Channel
	if err := ctx.BindJSON(&form); err != nil || len(form.Title) == 0 {
		base.BadRequestResponse(ctx, "invalid params")
		return
	}

	result, err := models.ChannelModel.Create(form.Title)
	if err != nil {
		base.ServerErrorResponse(ctx, err.Error())
		return
	}
	base.SuccessResponse(ctx, result)
}

func DeleteController(ctx *gin.Context) {
	channelId, err := parsers.ParserInt64(ctx.Param("channelId"))

	if err != nil {
		base.BadRequestResponse(ctx, err.Error())
		return
	}

	if err = models.ChannelModel.Delete(channelId); err != nil {
		base.ServerErrorResponse(ctx, err.Error())
		return
	}
	if err = models.MessageModel.Delete(channelId); err != nil {
		base.ServerErrorResponse(ctx, err.Error())
		return
	}
	base.SuccessResponse(ctx, channelId)
}

func UpdateController(ctx *gin.Context) {
	var form models.Channel

	channelId, err := parsers.ParserInt64(ctx.Param("channelId"))

	if err != nil {
		base.BadRequestResponse(ctx, err.Error())
		return
	}

	if err = ctx.BindJSON(&form); err != nil || len(form.Title) == 0 {
		base.BadRequestResponse(ctx, "invalid params")
		return
	}

	if err = models.ChannelModel.Update(channelId, form.Title); err != nil {
		base.ServerErrorResponse(ctx, err.Error())
		return
	}
	form.ID = channelId
	base.SuccessResponse(ctx, form)
}

func RetrieveController(ctx *gin.Context) {
	channelId, err := parsers.ParserInt64(ctx.Param("channelId"))
	if err != nil {
		base.BadRequestResponse(ctx, err.Error())
		return
	}
	obj := models.ChannelModel.Get(channelId)
	if obj.ID == 0 {
		base.SuccessResponse(ctx, nil)
		return
	}
	base.SuccessResponse(ctx, obj)
}
