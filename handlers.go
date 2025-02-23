package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KarlHavoc/aggregatorCLI/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("useage: %s <name>", cmd.Name)
	}
	name := cmd.Arguments[0]
	_, err := s.db.GetUser(context.Background(), sql.NullString{String: name, Valid: true})
	if err != nil {
		log.Fatal(err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %v", name)
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("please enter a name")
	}
	name := cmd.Arguments[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{uuid.New(), time.Now(), time.Now(), sql.NullString{String: name, Valid: true}})
	if err != nil {
		log.Fatal(err)
	}
	s.cfg.SetUser(name)
	fmt.Printf("user %s was created\n", name)
	log.Printf("user: %v\ncreated_at: %v\nupdated_at: %v\nname: %v\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name.String)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Println("failed to delete users table")
	}
	fmt.Println("deleted all users from db")
	os.Exit(0)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		if s.cfg.CurrentUserName == user.Name.String {
			fmt.Printf("%s (current)\n", user.Name.String)
			continue
		}
		fmt.Println(user.Name.String)
	}

	return nil
}
