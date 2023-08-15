package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/res"
	"github.com/pewpewnor/rss-aggregator/src/utils"
)

func (hc *HandlerContext) HandleCreateFeed(c *gin.Context) {
	user := utils.GetUserFromAuthMiddleware(c)

	var params struct {
		Name string `json:"name" binding:"required"`
		Url  string `json:"url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, res.SimpleErrorResponseFromError(
			"Invalid request body", err))
		return
	}

	feed, err := hc.DB.CreateFeed(c, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		c.JSON(400, res.SimpleErrorResponseFromError(
			"Cannot create feed in database", err))
		return
	}

	c.JSON(201, res.SuccessResponse(
		gin.H{"feed": utils.DBFeedToModelFeed(feed)},
		"Feed successfully created"))
}
