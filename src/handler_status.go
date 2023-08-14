package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pewpewnor/rss-aggregator/src/res"
)

func handleReady(c *gin.Context) {
	c.JSON(200, res.SimpleSuccessResponse("Server is ready"))
}
