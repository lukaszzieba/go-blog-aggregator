package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/lukaszzieba/go-blog-agregator/internal"
	"github.com/lukaszzieba/go-blog-agregator/internal/database"
)

func main() {
	c, err := internal.ReadConfig()
	if err != nil {
		fmt.Println("read config filed")
		panic(err)
	}
	db, err := sql.Open("postgres", c.Db_url)
	if err != nil {
		fmt.Println("open db failed")
		panic(err)
	}
	dbQueries := database.New(db)
	satate := internal.NewState(c, dbQueries)
	args, err := getArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	commands := internal.NewCommands()
	commands.Register("login", internal.HandlerLogin)
	commands.Register("register", internal.HandlerRegister)
	commands.Register("reset", internal.HandlerReset)
	commands.Register("users", internal.HandlerUsers)
	commands.Register("agg", internal.HandleAgg)
	commands.Register("addfeed", internal.MiddlewareLoggedIn(internal.HandleAddFeed))
	commands.Register("feeds", internal.HandlerFeeds)
	commands.Register("follow", internal.MiddlewareLoggedIn(internal.HandlerFeedFollow))
	commands.Register("following", internal.MiddlewareLoggedIn(internal.HandlerFeedFollowing))
	commands.Register("unfollow", internal.MiddlewareLoggedIn(internal.HandlerFeedUnfollow))
	commands.Register("browse", internal.MiddlewareLoggedIn(internal.HandleBrowse))

	err = commands.Run(satate, internal.Command{Name: args[0], Args: args[1:]})
	if err != nil {
		log.Fatal(err)
	}
}

func getArgs() ([]string, error) {
	args := os.Args[1:]
	if len(args) < 1 {
		return nil, fmt.Errorf("wrong parameter number")
	}

	return args, nil
}
