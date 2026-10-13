package main

import (
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("We need a time duration ")
	}
	fetchBetweenReqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v", cmd.arguments[0])

	ticker := time.NewTicker(fetchBetweenReqs)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

}
