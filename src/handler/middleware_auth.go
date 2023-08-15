package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pewpewnor/rss-aggregator/src/model"
	"github.com/pewpewnor/rss-aggregator/src/res"
	"github.com/pewpewnor/rss-aggregator/src/utils"
)

func (hc *HandlerContext) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey, err := utils.GetAPIKey(c)
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
			c.AbortWithStatusJSON(401, res.SimpleErrorResponse(
				"Authentication error", "User with API key not found"))
			return
		}

		c.Set("user", model.User(user))
		c.Next()
	}
}
