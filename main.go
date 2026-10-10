package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/JtorresBear/gator_project/internal/config"
	"github.com/JtorresBear/gator_project/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", cfg.Database)
	dbQueries := database.New(db)
	st8 := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		cmdMap: make(map[string]func(*state, command) error),
	}

	fillCommands(&cmds)

	userInputs := os.Args
	if len(userInputs) < 2 {
		log.Fatal("too few user inputs")
	}
	handlerName := userInputs[1]
	handlerArgs := userInputs[2:]
	cmdHandle := command{
		name:      handlerName,
		arguments: handlerArgs,
	}

	err = cmds.run(&st8, cmdHandle)
	if err != nil {
		log.Fatal(err)
	}
}

func fillCommands(c *commands) {
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", handlerFollowing)
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))
}
