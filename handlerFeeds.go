package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println("\nFeed Name: ", feed.FeedName)
		fmt.Println("Feed URL: ", feed.Url)
		fmt.Println("User Name: ", feed.UserName)
	}
	return nil
}
