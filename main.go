package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	internal "github.com/KarlHavoc/aggregatorCLI/internal/config"
	"github.com/KarlHavoc/aggregatorCLI/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *internal.Config
}

func main() {
	config_data, err := internal.ReadConfig()
	if err != nil {
		fmt.Print(err)
	}
	db, err := sql.Open("postgres", config_data.DbURL)
	if err != nil {
		log.Fatal("failed to load db")
	}
	dbQueries := database.New(db)

	programState := &state{
		cfg: &config_data,
		db:  dbQueries,
	}

	cmds := commands{
		command_map: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", cmds.aggregate)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
		return
	}

	cmd_name := os.Args[1]
	cmd_args := os.Args[2:]

	err = cmds.run(programState, command{Name: cmd_name, Arguments: cmd_args})
	if err != nil {
		log.Fatal(err)
	}
}
