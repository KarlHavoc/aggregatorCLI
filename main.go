package main

import (
	"context"
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
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}

func scrapeFeeds(s *state) error {
	feed_to_fetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	s.db.MarkFeedFetched(context.Background(), feed_to_fetch.ID)
	new_feed, err := fetchFeed(context.Background(), feed_to_fetch.Url)
	if err != nil {
		return err
	}
	for _, item := range new_feed.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}
