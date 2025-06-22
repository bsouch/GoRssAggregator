package main

import (
	"time"

	"github.com/bsouch/GoRssAggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json="api_key"`
}

func dboUserToUser(dboUser database.User) User {
	return User{
		ID:        dboUser.ID,
		CreatedAt: dboUser.CreatedAt,
		UpdatedAt: dboUser.UpdatedAt,
		Name:      dboUser.Name,
		ApiKey:    dboUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json="id"`
	CreatedAt time.Time `json="created_at"`
	UpdatedAt time.Time `json="updated_at"`
	Name      string    `json="name"`
	Url       string    `json="url"`
	UserID    uuid.UUID `json="user_id"`
}

func dboFeedToFeed(dboFeed database.Feed) Feed {
	return Feed{
		ID:        dboFeed.ID,
		CreatedAt: dboFeed.CreatedAt,
		UpdatedAt: dboFeed.UpdatedAt,
		Name:      dboFeed.Name,
		Url:       dboFeed.Url,
		UserID:    dboFeed.UserID,
	}
}

func dboFeedsToFeeds(dboFeed []database.Feed) []Feed {
	feeds := []Feed{}
	for i := 0; i < len(dboFeed); i++ {
		feeds = append(feeds, dboFeedToFeed(dboFeed[i]))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json="id"`
	CreatedAt time.Time `json="created_at"`
	UpdatedAt time.Time `json="updated_at"`
	UserID    uuid.UUID `json="user_id"`
	FeedID    uuid.UUID `json="feed_id"`
}

func dboFeedFollowToFeedFollow(dboFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dboFeedFollow.ID,
		CreatedAt: dboFeedFollow.CreatedAt,
		UpdatedAt: dboFeedFollow.UpdatedAt,
		UserID:    dboFeedFollow.UserID,
		FeedID:    dboFeedFollow.FeedID,
	}
}

func dboFeedFollowsToFeedFollows(dboFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for i := 0; i < len(dboFeedFollows); i++ {
		feedFollows = append(feedFollows, dboFeedFollowToFeedFollow(dboFeedFollows[i]))
	}
	return feedFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func dboPostToPost(dboPost database.Post) Post {
	var description *string
	if dboPost.Description.Valid {
		description = &dboPost.Description.String
	}

	return Post{
		ID:          dboPost.ID,
		CreatedAt:   dboPost.CreatedAt,
		UpdatedAt:   dboPost.UpdatedAt,
		Title:       dboPost.Title,
		Description: description,
		PublishedAt: dboPost.PublishedAt,
		Url:         dboPost.Url,
		FeedID:      dboPost.FeedID,
	}
}

func dboPostsToPosts(dboPosts []database.Post) []Post {
	posts := []Post{}
	for i := 0; i < len(dboPosts); i++ {
		posts = append(posts, dboPostToPost(dboPosts[i]))
	}
	return posts
}
