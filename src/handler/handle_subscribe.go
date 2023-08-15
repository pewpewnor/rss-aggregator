package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/model"
	"github.com/pewpewnor/rss-aggregator/src/res"
	"github.com/pewpewnor/rss-aggregator/src/utils"
)

func (hc *HandlerContext) HandleCreateSubscribe(c *gin.Context) {
	user := utils.GetUserFromAuthMiddleware(c)

	var params struct {
		FeedID uuid.UUID `json:"feed_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, res.SimpleErrorResponseFromError(
			"Invalid request body", err))
		return
	}

	subscribe, err := hc.DB.CreateSubscribe(
		c, database.CreateSubscribeParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			FeedID:    params.FeedID,
		})
	if err != nil {
		c.JSON(400, res.SimpleErrorResponseFromError(
			"Cannot create subscribe in database", err))
		return
	}

	c.JSON(201, res.SuccessResponse(
		gin.H{"subscribe": model.Subscribe(subscribe)},
		"Subscribe successfully created"))
}

func (hc *HandlerContext) HandleGetSubscribe(c *gin.Context) {
	user := utils.GetUserFromAuthMiddleware(c)

	subscribe, err := hc.DB.GetSubscribe(c, user.ID)
	if err != nil {
		c.JSON(400, res.SimpleErrorResponseFromError(
			"Cannot get subscribe from database", err))
		return
	}

	c.JSON(201, res.SuccessResponse(
		gin.H{"subscribe": utils.DBSubscribesToModelSubscribes(subscribe)},
		"Subscribe successfully found"))
}
