package utils

import (
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/model"
)

func DBUserToModelUser(dbUser database.User) model.User {
	return model.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}
