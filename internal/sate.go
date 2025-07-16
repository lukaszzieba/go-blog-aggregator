package internal

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lukaszzieba/go-blog-agregator/internal/database"
)

type State struct {
	db     *database.Queries
	Config *Config
}

func NewState(c *Config, db *database.Queries) *State {
	return &State{Config: c, db: db}
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login cmd requires name")
	}

	userName := cmd.Args[0]
	user, _ := s.db.GetUser(context.Background(), userName)
	if user == (database.User{}) {
		fmt.Printf("user with name %s don't exists\n", userName)
		os.Exit(1)
	}

	_, err := s.Config.SetUser(userName)
	if err != nil {
		return err
	}

	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("register cmd requires name")
	}

	userName := cmd.Args[0]

	if _, err := s.db.GetUser(context.Background(), userName); err == nil {
		fmt.Printf("user with name %s already exists\n", userName)
		os.Exit(1)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: userName})
	if err != nil {
		fmt.Println(err)
	}
	s.Config.SetUser(user.Name)
	return nil
}
