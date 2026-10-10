package main

import (
	"context"
	"errors"

	"github.com/JtorresBear/gator_project/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Use correct amount of arguments 1")
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	deleteParams := database.DeleteFeedFollowParams{}
	deleteParams.UserID = user.ID
	deleteParams.FeedID = feed.ID
	err = s.db.DeleteFeedFollow(context.Background(), deleteParams)
	if err != nil {
		return err
	}
	return nil
}
