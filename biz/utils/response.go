package utils

import (
	"github.com/gin-gonic/gin"
	"lutasam/GIN_LUTA/biz/bo"
)

func Response(c *gin.Context, code int, msg string, data interface{}) {
	resp := &bo.BaseResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(code, resp)
}
