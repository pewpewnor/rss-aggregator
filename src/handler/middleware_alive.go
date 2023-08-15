package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pewpewnor/rss-aggregator/src/res"
)

func (hc *HandlerContext) AliveMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := hc.DB.Alive(c)
		if err != nil {
			c.AbortWithStatusJSON(500, res.SimpleErrorResponseFromError(
				"Server inactive (dead)", err))
			return
		}

		c.Next()
	}
}
