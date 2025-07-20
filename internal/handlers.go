package internal

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/lukaszzieba/go-blog-agregator/internal/database"
)

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

	_, err := s.Config.SetUser(user)
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

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: userName})
	if err != nil {
		fmt.Println(err)
	}
	s.Config.SetUser(user)
	return nil
}

func HandlerUsers(s *State, cmd Command) error {
	data, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range data {
		fmt.Println(getUserLine(u, s.Config.Current_user))
	}

	return nil
}

func getUserLine(u database.User, currentUserName database.User) string {
	s := fmt.Sprint("*" + " " + u.Name)

	if u.Name == currentUserName.Name {
		s = fmt.Sprint(s + " " + "(current)")
	}

	return s
}

func HandlerReset(s *State, cmd Command) error {
	return s.db.DeleteAllUsers(context.Background())
}

func HandleAgg(s *State, cmd Command) error {
	res, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

func HandleAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		fmt.Println("add feed command rqeires 2 args: name, url")
		os.Exit(1)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name, Url: url, UserID: s.Config.Current_user.ID})
	if err != nil {
		return fmt.Errorf("create feed failed: %w", err)
	}

	fmt.Println(feed)
	return nil
}
