package main

import (
	"time"

	"github.com/20ritiksingh/RSSagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}
type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}
type post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func dbUserToUser(dbUser database.User) User {
	user := User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
	return user
}
func dbFeedToFeed(dbFeed database.Feed) Feed {
	feed := Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
	return feed
}
func dbFeedsToFeeds(dbFeed []database.Feed) []Feed {
	feeds := []Feed{}

	for _, feed := range dbFeed {
		feeds = append(feeds, dbFeedToFeed(feed))
	}

	return feeds
}
func dbFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	feedFollow := FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
	return feedFollow
}
func dbFeedFollowsToFeedFollows(dbFeed []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, feed := range dbFeed {
		feedFollows = append(feedFollows, dbFeedFollowToFeedFollow(feed))
	}

	return feedFollows
}

func dbPostToPost(dbposts database.Post) post {
	var description *string
	if dbposts.Description.Valid {
		description = &dbposts.Description.String
	}
	return post{
		ID:          dbposts.ID,
		CreatedAt:   dbposts.CreatedAt,
		UpdatedAt:   dbposts.UpdatedAt,
		Title:       dbposts.Title,
		Description: description,
		PublishedAt: dbposts.PublishedAt,
		Url:         dbposts.Url,
		FeedID:      dbposts.FeedID,
	}
}

func dbPostsToPosts(dbposts []database.Post) []post {
	posts := []post{}

	for _, post := range dbposts {
		posts = append(posts, dbPostToPost(post))
	}

	return posts
}
