package main

import "github.com/gin-gonic/gin"

type response struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, response{
			Message: "hi picpay",
		})
	})
	router.Run("localhost:8000")
}
