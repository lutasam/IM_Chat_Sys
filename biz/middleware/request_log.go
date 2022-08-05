package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/utils"
	"time"
)

func RequestDoTracerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceContext := utils.NewTrace()
		c.Set("startTime", time.Now())
		c.Set("trace", traceContext)
		c.Next()
	}
}
