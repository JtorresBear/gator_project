package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}
	RSSFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, item := range RSSFeed.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}
