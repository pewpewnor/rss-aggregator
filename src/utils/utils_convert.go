package utils

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

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

func UrlToFeed(url string) (model.RSSFeed, error) {
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
