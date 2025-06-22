package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bsouch/GoRssAggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(writer http.ResponseWriter, r *http.Request, dboUser database.User) {
	type parameters struct {
		Feed_ID uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	deserialiseErr := jsonDeserialise(r, &params)
	if deserialiseErr != nil {
		errorJsonResponse(writer, 400, fmt.Sprintf("Error parsing JSON: %v", deserialiseErr))
		return
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    dboUser.ID,
		FeedID:    params.Feed_ID,
	}

	dboFeedFollow, feedFollowError := apiCfg.DB.CreateFeedFollow(r.Context(), feedFollow)
	if feedFollowError != nil {
		errorJsonResponse(writer, 500, fmt.Sprintf("Error creating feed: %v", feedFollowError))
		return
	}

	jsonResponse(writer, 200, dboFeedFollowToFeedFollow(dboFeedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(writer http.ResponseWriter, r *http.Request, dboUser database.User) {
	type parameters struct {
		User_ID uuid.UUID `json:"user_id"`
	}
	params := parameters{}
	deserialiseErr := jsonDeserialise(r, &params)
	if deserialiseErr != nil {
		errorJsonResponse(writer, 400, fmt.Sprintf("Error parsing JSON: %v", deserialiseErr))
		return
	}

	dboFeedFollows, feedFollowsError := apiCfg.DB.GetFeedFollows(r.Context(), params.User_ID)
	if feedFollowsError != nil {
		errorJsonResponse(writer, 500, fmt.Sprintf("Error creating feed: %v", feedFollowsError))
		return
	}

	jsonResponse(writer, 200, dboFeedFollowsToFeedFollows(dboFeedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(writer http.ResponseWriter, r *http.Request, dboUser database.User) {
	feedFollowID := chi.URLParam(r, "feedFollowID")
	parsedFeedFollowID, feedFollowParseErr := uuid.Parse(feedFollowID)
	if feedFollowParseErr != nil {
		errorJsonResponse(writer, 400, fmt.Sprintf("Error pasring feed follow ID: %v", feedFollowParseErr))
	}

	delFeedFollows := database.DeleteFeedFollowParams{
		ID:     parsedFeedFollowID,
		UserID: dboUser.ID,
	}
	delFeedFollowsError := apiCfg.DB.DeleteFeedFollow(r.Context(), delFeedFollows)
	if delFeedFollowsError != nil {
		errorJsonResponse(writer, 500, fmt.Sprintf("Error deleting feed follow: %v", delFeedFollowsError))
		return
	}

	type deleteSuccess struct {
		Message string `json="message"`
	}
	deleteSuccessDto := deleteSuccess{Message: "Success!"}
	jsonResponse(writer, 200, deleteSuccessDto)
}
