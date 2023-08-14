package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pewpewnor/rss-aggregator/internal/auth"
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/res"
)

func (hc *HandlerContext) handleCreateUser(c *gin.Context) {
	var params struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, res.SimpleErrorResponseFromError(err))
		return
	}

	user, err := hc.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		c.JSON(500, res.SimpleErrorResponseFromError(err))
		return
	}

	c.JSON(201, res.SuccessResponse(gin.H{"user": dbUserToUser(user)}, "User successfully created"))
}

func (hc *HandlerContext) handleGetUser(c *gin.Context) {
	apiKey, err := auth.GetAPIKey(c)
	if err != nil {
		errorResponse, ok := err.(res.ErrorResponseData)
		if !ok {
			c.JSON(401, err.Error())
			return
		}

		c.JSON(401, errorResponse)
		return
	}

	user, err := hc.DB.GetUserByAPIKey(c, apiKey)
	if err != nil {
		c.JSON(401, res.SimpleErrorResponse("Authentication error", "User with API key not found"))
		return
	}

	c.JSON(200, res.SuccessResponse(gin.H{"user": dbUserToUser(user)}, "User successfully found"))
}
