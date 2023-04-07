package base

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BadRequestResponse(ctx *gin.Context, msg string) {
	ctx.JSON(400, Response{
		Code:    400,
		Message: msg,
		Data:    nil,
	})
}

func ServerErrorResponse(ctx *gin.Context, msg string) {
	ctx.JSON(500, Response{
		Code:    500,
		Message: msg,
		Data:    nil,
	})
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}
