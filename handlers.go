package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("useage: %s <name>", cmd.Name)
	}
	name := cmd.Arguments[0]
	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error setting user name: %w", err)
	}
	fmt.Println("User has been set")
	return nil
}
