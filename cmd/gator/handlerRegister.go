package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/JtorresBear/gator_project/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("Username is required")
	}
	new_user := database.CreateUserParams{}
	new_user.ID = uuid.New()
	time := time.Now()
	new_user.CreatedAt = time
	new_user.UpdatedAt = time
	new_user.Name = cmd.arguments[0]

	user, err := s.db.CreateUser(context.Background(), new_user)

	pqErr := new(pq.Error)
	if errors.As(err, &pqErr) {
		if pqErr.Code == "23505" {
			log.Fatal("Name already exists")
		}
	}
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Println("User was Created")
	log.Println("User ID: ", user.ID)
	log.Println("Created at: ", user.CreatedAt)
	log.Println("Updated at: ", user.UpdatedAt)
	log.Println("Name: ", user.Name)
	return nil
}
