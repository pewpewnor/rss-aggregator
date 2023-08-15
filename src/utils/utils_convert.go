package utils

import (
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/model"
)

func DBSubscribesToModelSubscribes(dbSubscribes []database.Subscribe) []model.Subscribe {
	subscribes := []model.Subscribe{}
	for _, dbSubscribe := range dbSubscribes {
		subscribes = append(subscribes, model.Subscribe(dbSubscribe))
	}
	return subscribes
}

// func DBUserToModelUser(dbUser database.User) model.User {
// 	return model.User{
// 		ID:        dbUser.ID,
// 		CreatedAt: dbUser.CreatedAt,
// 		UpdatedAt: dbUser.UpdatedAt,
// 		Name:      dbUser.Name,
// 		ApiKey:    dbUser.ApiKey,
// 	}
// }

// func DBFeedToModelFeed(dbFeed database.Feed) model.Feed {
// 	return model.Feed{
// 		ID:        dbFeed.ID,
// 		CreatedAt: dbFeed.CreatedAt,
// 		UpdatedAt: dbFeed.UpdatedAt,
// 		Name:      dbFeed.Name,
// 		URL:       dbFeed.Url,
// 		UserID:    dbFeed.UserID,
// 	}
// }

// func DBSubscribeToModelSubscribe(dbSubscribe database.Subscribe) model.Subscribe {
// 	return model.Subscribe{
// 		ID:        dbSubscribe.ID,
// 		CreatedAt: dbSubscribe.CreatedAt,
// 		UpdatedAt: dbSubscribe.UpdatedAt,
// 		UserID:    dbSubscribe.UserID,
// 		FeedID:    dbSubscribe.FeedID,
// 	}
// }
