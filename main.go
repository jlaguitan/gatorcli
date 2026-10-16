package main

import (
	"gatorcli/internal/config"
	"log"
	"os"
)

// state will give handlers access to the application state
type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
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
