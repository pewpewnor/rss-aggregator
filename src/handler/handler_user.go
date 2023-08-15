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

func (hc *HandlerContext) HandleCreateUser(c *gin.Context) {
	var params struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, res.SimpleErrorResponseFromError(
			"Invalid request body", err))
		return
	}

	user, err := hc.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		c.JSON(500, res.SimpleErrorResponseFromError(
			"Cannot create user in database", err))
		return
	}

	c.JSON(201, res.SuccessResponse(
		gin.H{"user": utils.DBUserToModelUser(user)}, "User successfully created"))
}

func (hc *HandlerContext) HandleGetUser(c *gin.Context) {
	user, _ := c.MustGet("user").(model.User)

	c.JSON(200, res.SuccessResponse(
		gin.H{"user": user}, "User successfully found"))
}
