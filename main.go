package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", )
	r.Run(":8080")
}

func PingHandler(c *gin.Context){
	res := Response{Message: "pong"}
	c.JSON(http.StatusOK, res)
}

type Response struct{
	Message string `json:"message"`
}