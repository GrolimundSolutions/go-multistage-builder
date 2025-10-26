package main

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/GrolimundSolutions/go-multistage-builder/api/v1"
)

func main() {
	r := gin.Default()
	r.GET("/v1/user/:id", v1.GetUser)
	r.POST("/v1/user", v1.CreateUser)
	r.Run()
}