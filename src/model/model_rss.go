package model

type RSSFeed struct {
	Channel struct {
		Title       string        `xml:"title"`
		Link        string        `xml:"link"`
		Description string        `xml:"description"`
		Language    string        `xml:"language"`
		Item        []RSSFeedItem `xml:"item"`
	} `xml:"channel"`
}

type RSSFeedItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
