package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bsouch/GoRssAggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(writer http.ResponseWriter, r *http.Request, dboUser database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"Url"`
	}
	params := parameters{}
	deserialiseErr := jsonDeserialise(r, &params)
	if deserialiseErr != nil {
		errorJsonResponse(writer, 400, fmt.Sprintf("Error parsing JSON: %v", deserialiseErr))
		return
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    dboUser.ID,
	}

	dboFeed, feedError := apiCfg.DB.CreateFeed(r.Context(), feed)
	if feedError != nil {
		errorJsonResponse(writer, 500, fmt.Sprintf("Error creating feed: %v", feedError))
		return
	}

	jsonResponse(writer, 200, dboFeedToFeed(dboFeed))
}

func (apiCfg *apiConfig) handlerGetFeeds(writer http.ResponseWriter, r *http.Request) {
	dboFeeds, feedsErr := apiCfg.DB.GetFeeds(r.Context())
	if feedsErr != nil {
		errorJsonResponse(writer, 500, fmt.Sprintf("Error getting feeds: %v", feedsErr))
		return
	}

	feeds := dboFeedsToFeeds(dboFeeds)
	jsonResponse(writer, 200, feeds)
}
