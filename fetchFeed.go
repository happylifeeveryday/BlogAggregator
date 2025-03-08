package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()
	var RSS RSSFeed
	xml.Unmarshal(byteData, &RSS)
	RSS.Channel.Title = html.UnescapeString(RSS.Channel.Title)
	RSS.Channel.Description = html.UnescapeString(RSS.Channel.Description)
	for index, _ := range RSS.Channel.Item {
		RSS.Channel.Item[index].Title = html.UnescapeString(RSS.Channel.Item[index].Title)
		RSS.Channel.Item[index].Description = html.UnescapeString(RSS.Channel.Item[index].Description)
	}
	return &RSS, nil
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
