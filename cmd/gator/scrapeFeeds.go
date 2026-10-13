package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/JtorresBear/gator_project/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	RSSFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	for _, item := range RSSFeed.Channel.Item {
		postParams := database.Create_PostParams{}
		timeP := time.Now()
		postParams.ID = uuid.New()
		postParams.CreatedAt = timeP
		postParams.UpdatedAt = timeP
		postParams.Title = sql.NullString{
			String: item.Title,
			Valid:  item.Title != "",
		}
		postParams.Url = item.Link
		postParams.Description = sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}
		postParams.PublishedAt = sql.NullTime{}
		parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err == nil {
			postParams.PublishedAt.Time = parsedTime
			postParams.PublishedAt.Valid = true
		}
		postParams.FeedID = feed.ID
		_, err = s.db.Create_Post(context.Background(), postParams)
		if err != nil {
			if strings.Contains(err.Error(), "posts_url_key") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
		}
	}
	return nil
}
