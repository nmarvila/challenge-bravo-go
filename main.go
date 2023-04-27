package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func rootRoute(c *gin.Context) {
	c.String(http.StatusOK, "Currency converter")
}

func main() {
	router := gin.Default()
	router.GET("/", rootRoute)

	router.Run("localhost:8080")
}