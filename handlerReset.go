package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		log.Fatalf("There was an err: %v", err)
		return err
	}
	fmt.Println("users were reset")
	return nil
}
