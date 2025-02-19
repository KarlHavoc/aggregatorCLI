package main

import (
	"fmt"
	"os"

	"github.com/KarlHavoc/gatorCLI/main/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	config_data, err := config.ReadConfig()
	if err != nil {
		fmt.Print(err)
	}
	arguments := os.Args

	new_command := command{
		Name:      arguments[1],
		Arguments: arguments[2:],
	}
	fmt.Println(new_command.Arguments)
	fmt.Println("searching for json config file")

	fmt.Println(config_data)

}
