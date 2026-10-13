package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/JtorresBear/gator_project/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.arguments) > 0 {
		var err error
		limit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return err
		}
	}
	params := database.Get_Posts_for_UserParams{}
	if limit < 1 {
		return errors.New("You need at least 1.")
	}
	params.Limit = int32(limit)
	params.UserID = user.ID
	posts, err := s.db.Get_Posts_for_User(context.Background(), params)
	if err != nil {
		return err
	}

	if len(posts) == 0 {
		fmt.Println("No posts found")
		return nil
	}
	for _, post := range posts {
		fmt.Println()
		fmt.Println(post.Title.String)
		fmt.Println()
		fmt.Println(post.Description.String)
		fmt.Println()
	}
	return nil
}
