package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", myGetEvents)

	server.Run(":8080")

}

func myGetEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})

}
