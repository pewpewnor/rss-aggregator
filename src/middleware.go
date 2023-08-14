package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pewpewnor/rss-aggregator/internal/auth"
	"github.com/pewpewnor/rss-aggregator/src/res"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (hc *HandlerContext) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey, err := auth.GetAPIKey(c)
		if err != nil {
			errorResponse, ok := err.(res.ErrorResponseData)
			if !ok {
				c.AbortWithStatusJSON(401, err.Error())
				return
			}

			c.AbortWithStatusJSON(401, errorResponse)
			return
		}

		user, err := hc.DB.GetUserByAPIKey(c, apiKey)
		if err != nil {
			c.AbortWithStatusJSON(401, res.SimpleErrorResponse("Authentication error", "User with API key not found"))
			return
		}

		c.Set("user", dbUserToModelUser(user))
		c.Next()
	}
}
