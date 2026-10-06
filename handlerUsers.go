package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name.String == s.cfg.Current_user_name {
			fmt.Printf("* %v (current)\n", user.Name.String)
		} else {
			fmt.Printf("* %v \n", user.Name.String)
		}
	}

	return nil
}
