package main

import (
	"fmt"
	"net/http"

	"github.com/20ritiksingh/RSSagg/internal/auth"
	"github.com/20ritiksingh/RSSagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("auth error : %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("couldn't get user : %v", err))
			return
		}

		handler(w, r, user)
	}
}
