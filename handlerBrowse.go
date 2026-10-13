package main

import (
	"context"
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
	params.Limit = int32(limit)
	params.UserID = user.ID
	posts, err := s.db.Get_Posts_for_User(context.Background(), params)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post)
	}
	return nil
}
