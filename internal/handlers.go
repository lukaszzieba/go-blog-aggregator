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
		return fmt.Errorf("follow cmd requires url")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name, Url: url, UserID: s.Config.Current_user.ID})
	if err != nil {
		return fmt.Errorf("create feed failed: %w", err)
	}
	feedFlow := database.CreateFeedFollowParams{
		UserID:    s.Config.Current_user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFlow)
	if err != nil {
		return fmt.Errorf("create feed follow failed: %w", err)
	}

	fmt.Println(feed)
	return nil
}

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.db.GetFeedsWithUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %s (%s) - created by %s\n", feed.Name, feed.Url, feed.UserName)
	}

	return nil
}

func HandlerFeedFollow(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("follow cmd requires url")
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	followFeed := database.CreateFeedFollowParams{UserID: s.Config.Current_user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now(), FeedID: feed.ID}
	res, err := s.db.CreateFeedFollow(context.Background(), followFeed)
	if err != nil {
		return err
	}
	printFeedFollow(res)

	return nil
}

func HandlerFeedFollowing(s *State, cmd Command) error {
	feeds, err := s.db.GetFeedsForUser(context.Background(), s.Config.Current_user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("User name			: %s\n", s.Config.Current_user.Name)
	fmt.Println("Feed names:")
	for _, f := range feeds {
		fmt.Printf("- %s\n", f.String)
	}

	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
	fmt.Printf("User name:				%s\n", feedFollow.UserName.String)
	fmt.Printf("Feed name:				%s\n", feedFollow.FeedName.String)
}
