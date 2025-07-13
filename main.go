package main

import (
	"os"

	"github.com/lukaszzieba/go-blog-agregator/internal"
)

func main() {
	c, err := internal.ReadConfig()
	if err != nil {
		panic(err)
	}
	args := os.Args[1:]
	if len(args) < 2 {
		os.Exit(1)
	}
	satate := internal.NewState(c)
	commands := internal.NewCommands()
	commands.Register("login", internal.HandlerLogin)
	commands.Run(satate, internal.Command{Name: args[0], Args: args[1:]})
}
