package utils

import (
	"database/sql"

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

func DBPostsToModelPosts(dbPosts []database.Post) []model.Post {
	posts := []model.Post{}
	for _, post := range dbPosts {
		posts = append(posts, DBPostToModelPost(post))
	}
	return posts
}

func DBPostToModelPost(dbPost database.Post) model.Post {
	return model.Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Url:         dbPost.Url,
		Title:       dbPost.Title,
		Description: NullStringToPointerString(dbPost.Description),
		PublishedAt: dbPost.PublishedAt,
		FeedID:      dbPost.FeedID,
	}
}

func NullStringToPointerString(nullString sql.NullString) string {
	if nullString.Valid {
		return nullString.String
	} else {
		return ""
	}
}
