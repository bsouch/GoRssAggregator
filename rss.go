package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type Channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Language    string    `xml:"language"`
	Item        []RSSItem `xml:"item"`
}

type RSSFeed struct {
	Channel Channel `xml:"channel"`
}

func UrlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	RSSFeed := RSSFeed{}
	xml.Unmarshal(data, &RSSFeed)
	return RSSFeed, nil
}
