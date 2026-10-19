package main

//sudo sed -i 's/trust/peer/' /etc/postgresql/*/main/pg_hba.conf && sudo service postgresql restart
//psql "postgres://postgres:postgres@localhost:5432/gator"
import (
	"context"
	"errors"
	"fmt"
	"gatorcli/internal/database"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	regCommand, ok := c.registeredCommands[cmd.name]
	if ok == false {
		return errors.New("command is not registered")
	}

	return regCommand(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("username is required")
	}

	userName := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		log.Printf("User:%v does not exist", userName)
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return errors.New("command failed: user cannot be set.")
	}
	fmt.Printf("User: %v has been logged in successfully\n", userName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("username is required")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		log.Print(err)
		log.Println("error: username already exists")
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return errors.New("command failed: username cannot be registered")
	}
	fmt.Printf("User: %v has been successfully registered\n", user.Name)
	log.Printf("%+v\n", user)
	return nil
}
