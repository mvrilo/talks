package main

import (
	"github.com/gin-gonic/gin"
	pb "github.com/mvrilo/talks/go-protobuf-microservices/generated"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		json := &pb.MessageResponse{
			Message: "hi picpay",
		}
		ctx.JSON(200, json)
	})
	router.Run("localhost:8000")
}
