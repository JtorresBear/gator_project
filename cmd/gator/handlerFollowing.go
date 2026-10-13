package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	currentUser := s.cfg.Current_user_name
	fmt.Println(currentUser, " is following: ")
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser)
	if err != nil {
		return err
	}
	if len(feeds) == 0 {
		fmt.Println("not a damn thing")
		return nil
	}
	for _, feed := range feeds {
		fmt.Println(feed.Name)
	}
	return nil
}
