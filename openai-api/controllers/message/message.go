package message

import (
	"bytes"
	"encoding/json"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/config"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/controllers/base"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/models"
	"github.com/YuZongYangHi/chatgpt-proxy/openai-api/pkg/util/parsers"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func ListMessage(ctx *gin.Context) {
	channelId, err := parsers.ParserInt64(ctx.Param("channelId"))

	if err != nil {
		base.BadRequestResponse(ctx, err.Error())
		return
	}

	result := models.MessageModel.List(channelId)
	base.SuccessResponse(ctx, result)
}

func CreateMessage(ctx *gin.Context) {
	channelId, err := parsers.ParserInt64(ctx.Param("channelId"))

	if err != nil {
		base.BadRequestResponse(ctx, err.Error())
		return
	}

	channel := models.ChannelModel.Get(channelId)
	if channel.ID == 0 {
		base.BadRequestResponse(ctx, "invalid channelId")
		return
	}

	var form models.Message

	if err = ctx.BindJSON(&form); err != nil || !form.IsValid() {
		base.BadRequestResponse(ctx, "params error")
		return
	}

	userMessage := &models.Message{
		ChannelId:  channelId,
		Content:    form.Content,
		DialogType: 1,
		CreatedAt:  time.Now(),
	}

	userMessage, err = models.MessageModel.Add(userMessage)
	if err != nil {
		base.ServerErrorResponse(ctx, err.Error())
		return
	}

	data := map[string]string{
		"content": form.Content,
		"token":   config.AppConfig().Proxy.Token,
	}

	jsonStr, _ := json.Marshal(data)

	tr := &http.Transport{
		DisableCompression: true,
	}
	request := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", config.AppConfig().Proxy.URL, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")

	resp, _ := request.Do(req)
	defer resp.Body.Close()

	var result = ""
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		result = string(body)
	} else {
		result = "An unknown problem occurred when obtaining the result"
	}

	robotMessage := &models.Message{
		ChannelId:  channelId,
		Content:    result,
		DialogType: 0,
		CreatedAt:  time.Now(),
	}

	robotMessage, err = models.MessageModel.Add(robotMessage)
	if err != nil {
		base.ServerErrorResponse(ctx, err.Error())
		return
	}

	base.SuccessResponse(ctx, robotMessage)
}
