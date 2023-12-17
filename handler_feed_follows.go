package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/20ritiksingh/RSSagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing json: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error creating feed : %v", err))
		return
	}

	respondWithJOSN(w, 201, dbFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't get your feeds: %v", err))
		return
	}
	respondWithJOSN(w, 201, dbFeedFollowsToFeedFollows(feedFollow))
}
func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't parse feed_id: %v", err))
		return
	}

	deleteFeedFollowParams := database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), deleteFeedFollowParams)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("couldn't delete feed_follow: %v", err))
		return
	}
	respondWithJOSN(w, 200, struct{}{})
}
