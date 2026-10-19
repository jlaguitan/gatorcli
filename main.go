package main

import (
	"database/sql"
	"log"
	"os"

	"gatorcli/internal/config"
	"gatorcli/internal/database"

	_ "github.com/lib/pq"
)

// state will give handlers access to the application state
type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	userInput := os.Args

	if len(userInput) < 2 {
		log.Fatal("invalid input")
	}

	commandName := os.Args[1]
	commandArgument := os.Args[2:]
	err = cmds.run(s, command{name: commandName, args: commandArgument})
	if err != nil {
		log.Fatal(err)
	}
}
