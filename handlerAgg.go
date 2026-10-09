package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerAgg(_ *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return errors.New("There was a problem fetching the feed")
	}
	fmt.Println(feed)
	return nil
}
