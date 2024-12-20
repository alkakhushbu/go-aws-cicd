package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.Run(":80")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello ping v1")
}
