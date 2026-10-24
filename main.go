package main

//sudo sed -i 's/trust/peer/' /etc/postgresql/*/main/pg_hba.conf && sudo service postgresql restart
//psql "postgres://postgres:postgres@localhost:5432/gator" //testing connection string
//goose postgres '<connection_string>' up
//sudo service postgresql start
//sudo service postgresql status
//psql -U postgres -d gator
import (
	"context"
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
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser := s.cfg.CurrentUserName
		user, err := s.db.GetUser(context.Background(), currentUser)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
