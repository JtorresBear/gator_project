package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Username is required")
	}
	user := sql.NullString{
		String: cmd.arguments[0],
		Valid:  true,
	}
	_, err := s.db.GetUser(context.Background(), user)

	if errors.Is(err, sql.ErrNoRows) {
		log.Fatal("No Such User")
	}

	err = s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}
