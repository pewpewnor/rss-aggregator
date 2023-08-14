package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pewpewnor/rss-aggregator/src/response"
)

func handleReady(c *gin.Context) {
	c.JSON(200, response.GenerateSimpleSuccessResponse("Server is ready"))
}
