package main

import (
	"log"
	"net/http"

	"github.com/20ritiksingh/RSSagg/internal/database"
)

func (apiCfg *apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		log.Printf("failed to get posts : %v", err)
		return
	}
	respondWithJOSN(w, 200, dbPostsToPosts(posts))
}
