package main

import (
	"errors"
	"fmt"
	"time"
)

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	command_map map[string]func(*state, command) error
}

func (c commands) register(name string, f func(*state, command) error) {
	c.command_map[name] = f
}

func (c commands) run(s *state, cmd command) error {
	com, ok := c.command_map[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return com(s, cmd)

}

func (c commands) aggregate(s *state, cmd command) error {
	time_between_reqs := cmd.Arguments[0]
	duration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", duration)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		fmt.Print("\n\n")
		fmt.Println("Scraping feeds...")
		fmt.Print("\n\n")
		scrapeFeeds(s)
		fmt.Print("\n\n")
	}

}
