package main

import (
	"fmt"
	"net/http"

	"github.com/bsouch/GoRssAggregator/internal/database"
	"github.com/bsouch/GoRssAggregator/internal/database/auth"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, authError := auth.GetAPIKey(r.Header)
		if authError != nil {
			errorJsonResponse(w, 403, fmt.Sprintf("Auth error: %v", authError))
			return
		}

		dboUser, userErr := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if userErr != nil {
			errorJsonResponse(w, 500, fmt.Sprintf("Couldn't get user: %v", userErr))
			return
		}

		handler(w, r, dboUser)
	}
}
