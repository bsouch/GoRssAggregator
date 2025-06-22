package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bsouch/GoRssAggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(writer http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	err := jsonDeserialise(r, params)
	if err != nil {
		errorJsonResponse(writer, 400, fmt.Sprintf("Error parsing Json: %v", err))
		return
	}

	createUserDbo := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	}

	dboUser, err := apiCfg.DB.CreateUser(r.Context(), createUserDbo)
	if err != nil {
		errorJsonResponse(writer, 500, fmt.Sprintf("Could not creat user: %v", err))
		return
	}

	jsonResponse(writer, 201, dboUserToUser(dboUser))
}

func (apiCfg *apiConfig) handlerGetUser(writer http.ResponseWriter, r *http.Request, dboUser database.User) {
	jsonResponse(writer, 200, dboUserToUser(dboUser))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(writer http.ResponseWriter, r *http.Request, dboUser database.User) {
	params := database.GetPostsForUserParams{
		UserID: dboUser.ID,
		Limit:  10,
	}
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), params)
	if err != nil {
		log.Printf("Error fetching posts for user: %v", err)
	}

	jsonResponse(writer, 200, dboPostsToPosts(posts))
}
