package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/response"
)

func (hc *HandlerContext) handleCreateUser(c *gin.Context) {
	var params struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(400, response.SimpleErrorResponseFromError(err))
		return
	}

	user, err := hc.DB.CreateUser(c, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		c.JSON(500, response.SimpleErrorResponseFromError(err))
		return
	}

	c.JSON(200, response.SuccessResponse(gin.H{"user": dbUserToUser(user)}, "User successfully created"))
}

func (hc *HandlerContext) handleGetUser(c *gin.Context) {

}
