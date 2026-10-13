package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Username is required")
	}

	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])

	if errors.Is(err, sql.ErrNoRows) {
		return errors.New("No Such User")
	}
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}
