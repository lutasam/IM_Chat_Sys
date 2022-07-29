package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/bo"
	"github.com/lutasam/chat/biz/common"
	"github.com/lutasam/chat/biz/utils"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			utils.Response(c, common.CLIENTERRORCODE, common.CLIENTERRORMSG, &bo.BaseResponse{
				Code: common.USERNOTLOGIN.Code(),
				Msg:  common.USERNOTLOGIN.Error(),
				Data: nil,
			})
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer") {
			token = strings.Split(token, "")[1]
		}
		jwtStruct, err := utils.ParseJWTToken(token)
		if err != nil {
			utils.Response(c, common.CLIENTERRORCODE, common.CLIENTERRORMSG, &bo.BaseResponse{
				Code: common.EXCEEDTIMELIMIT.Code(),
				Msg:  common.EXCEEDTIMELIMIT.Error(),
				Data: nil,
			})
			c.Abort()
			return
		}
		c.Set("jwtStruct", jwtStruct)
		c.Next()
	}
}
