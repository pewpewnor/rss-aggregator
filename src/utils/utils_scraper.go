package utils

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pewpewnor/rss-aggregator/internal/database"
	"github.com/pewpewnor/rss-aggregator/src/logmsg"
	"github.com/pewpewnor/rss-aggregator/src/model"
)

func StartScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Println(logmsg.Infof(
		"scraping on %v goroutines every %v duration",
		concurrency, timeBetweenRequest))

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(), int32(concurrency))
		if err != nil {
			log.Println(logmsg.Warn("error getting next feeds from db:", err))
			continue
		}

		log.Println(logmsg.Infof("scraping %v RSS feeds", len(feeds)))

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.UpdateLastFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println(logmsg.Warn(
			"error marking feed as fetched (feed id does not exist):", err))
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println(logmsg.Warn("error fetching feed from URL ", feed.Url))
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		// TODO: add more robust date parsing logic
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println(logmsg.Warn(
				"Could not parse date, need more robust date parsing logic"))
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Url:         item.Link,
			Title:       item.Title,
			Description: description,
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				continue
			}
			log.Print(logmsg.Warn("failed to create post: ", err))
		}
	}

	log.Println(logmsg.Successf(
		"%v posts fetched from feed '%v'",
		len(rssFeed.Channel.Item), feed.Name))
}

func urlToFeed(url string) (model.RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return model.RSSFeed{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.RSSFeed{}, err
	}

	rssFeed := model.RSSFeed{}
	err = xml.Unmarshal(body, &rssFeed)
	if err != nil {
		return model.RSSFeed{}, err
	}

	return rssFeed, nil
}
