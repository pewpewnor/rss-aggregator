package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pewpewnor/rss-aggregator/src/res"
)

// Header format is the following
// Authorization: Apikey <token>

func GetAPIKey(c *gin.Context) (string, error) {
	value := c.Request.Header["Authorization"][0]
	if value == "" {
		err := res.SimpleErrorResponse("Authentication error", "No authorization header or its value is not given")
		return "", err
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		err := res.SimpleErrorResponse("Authentication error", "Malformed authorization header value")
		err.AddValidation(res.ErrorResponseValidation{
			Field:   "Authorization header",
			Message: "Expected exactly 2 values",
		})

		return "", err
	}
	if values[0] != "Apikey" {
		err := res.SimpleErrorResponse("Authentication error", "Malformed authorization header value")
		err.AddValidation(res.ErrorResponseValidation{
			Field:   "Authorization header",
			Message: "First value must be 'Apikey'",
		})

		return "", err
	}

	return values[1], nil
}