package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/JtorresBear/gator_project/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return errors.New("You need a name and url")
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return err
	}
	feedParams := database.CreateFeedParams{}
	feedParams.ID = uuid.New()
	timeF := time.Now()
	feedParams.CreatedAt = timeF
	feedParams.UpdatedAt = timeF
	feedParams.Name = cmd.arguments[0]
	feedParams.Url = cmd.arguments[1]
	feedParams.UserID = user.ID
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	followParams := database.CreateFeedFollowParams{}
	followParams.ID = uuid.New()
	followParams.CreatedAt = timeF
	followParams.UpdatedAt = timeF
	followParams.UserID = user.ID
	followParams.FeedID = feed.ID
	_, err = s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return err
	}
	fmt.Println("ID:", feed.ID)
	fmt.Println("Created AT: ", feed.CreatedAt)
	fmt.Println("Updated At: ", feed.UpdatedAt)
	fmt.Println("Name: ", feed.Name)
	fmt.Println("url: ", feed.Url)
	fmt.Println("UserID: ", feed.UserID)
	return nil
}
