package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/bsouch/GoRssAggregator/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, concurrency int, timeBetweeRequest time.Duration) {
	log.Printf("Scraping on v% GoRoutines every %s duration", concurrency, timeBetweeRequest)

	ticker := time.NewTicker(timeBetweeRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurrency))
		if err != nil {
			fmt.Printf("Error fetching feeds from DB: %v", err)
			continue
		}

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

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Failed to update feed as fetched: %v", err)
		return
	}

	rssFeed, err := UrlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{
			String: item.Description,
			Valid:  (item.Description != ""),
		}

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Unable to parse item publish date: %v", item.PubDate)
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Description: description,
			PublishedAt: pubDate,
			Url:         item.Link,
			FeedID:      feed.ID,
		}

		_, err = db.CreatePost(context.Background(), params)
		if err != nil && !strings.Contains(err.Error(), "duplicate key") {
			log.Printf("Unable to post to the DB: %v", err)
		}
	}
	log.Printf("Feed %s fetched, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
