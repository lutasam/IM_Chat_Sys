package dal

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/chat/biz/model"
	"sync"
)

type DemoDal struct{}

var (
	demoDal     *DemoDal
	demoDalOnce sync.Once
)

func GetDemoDal() *DemoDal {
	demoDalOnce.Do(func() {
		demoDal = &DemoDal{}
	})
	return demoDal
}

func (ins *DemoDal) Ping(c *gin.Context) (string, error) {
	return "pong", nil
}

func (ins *DemoDal) Hello(c *gin.Context) (*model.Hello, error) {
	return &model.Hello{
		Hello:  "hello",
		Author: "lutasam",
	}, nil
}
