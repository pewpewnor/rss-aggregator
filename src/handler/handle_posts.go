package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/logmsg"
	"github.com/pewpewnor/rss-aggregator/src/res"
	"github.com/pewpewnor/rss-aggregator/src/utils"
)

func (hc *HandlerContext) HandleGetNewestPostsForUser(c *gin.Context) {
	user := utils.GetUserFromAuthMiddleware(c)

	posts, err := hc.DB.GetNewestPostsForUser(c, database.GetNewestPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		c.JSON(500, res.SimpleErrorResponseFromError(
			"Get posts error", err))
		return
	}

	log.Print(logmsg.Success("posts gotten", len(posts)))

	c.JSON(200, res.SuccessResponse(
		gin.H{"posts": utils.DBPostsToModelPosts(posts)},
		"Posts successfully found"))
}
