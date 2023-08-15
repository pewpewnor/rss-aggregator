package utils

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pewpewnor/rss-aggregator/src/model"
	"github.com/pewpewnor/rss-aggregator/src/res"
)

// Header format is the following
// Authorization: Apikey <token>
func GetAPIKey(c *gin.Context) (string, error) {
	value := c.Request.Header["Authorization"][0]
	if value == "" {
		err := res.SimpleErrorResponse("Authentication error", "no authorization header or its value is not given")
		return "", err
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		err := res.SimpleErrorResponse("Authentication error", "malformed authorization header value")
		err.AddValidation(res.ErrorResponseValidation{
			Field:   "Authorization header",
			Message: "Expected exactly 2 values",
		})

		return "", err
	}
	if values[0] != "Apikey" {
		err := res.SimpleErrorResponse("Authentication error", "malformed authorization header value")
		err.AddValidation(res.ErrorResponseValidation{
			Field:   "Authorization header",
			Message: "First value must be 'Apikey'",
		})

		return "", err
	}

	return values[1], nil
}

func GetUserFromAuthMiddleware(c *gin.Context) model.User {
	anyUser, exists := c.Get("user")
	if !exists {
		if os.Getenv("PRODUCTION") == "true" {
			return model.User{}
		}
		panic("get 'user' does not exist (is not set from any middleware)")
	}
	user, ok := anyUser.(model.User)
	if !ok {
		if os.Getenv("PRODUCTION") == "true" {
			return model.User{}
		}
		panic("type assertion to convert to model.User failed")
	}

	return user
}
