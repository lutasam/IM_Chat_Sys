package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	InitRouterAndMiddleware(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
