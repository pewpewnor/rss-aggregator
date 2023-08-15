package handler

import "github.com/pewpewnor/rss-aggregator/internal/database"

type HandlerContext struct {
	DB *database.Queries
}
