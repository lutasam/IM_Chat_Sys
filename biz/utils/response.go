package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
)

func Response(c *gin.Context, code int, msg string, data interface{}) {
	resp := &bo.BaseResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(code, resp)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	Response(c, common.STATUSOKCODE, common.STATUSOKMSG, data)
}

func ResponseError(c *gin.Context, err common.Error) {
	Response(c, err.Code(), err.Error(), nil)
}
