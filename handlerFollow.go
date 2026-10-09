package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JtorresBear/gator_project/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Need One url argument")
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return err
	}
	followParams := database.CreateFeedFollowParams{}
	followParams.ID = uuid.New()
	timeN := time.Now()
	followParams.CreatedAt = timeN
	followParams.UpdatedAt = timeN
	followParams.UserID = user.ID
	followParams.FeedID = feed.ID
	row, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return err
	}
	fmt.Println("Feed: ", row.FeedName)
	fmt.Println("Followed by: ", row.UserName)
	return nil
}
