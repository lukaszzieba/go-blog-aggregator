package main

import (
	"fmt"

	"github.com/lukaszzieba/go-blog-agregator/internal/config"
)

func main() {
	c, err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}
	c2, err := c.SetUser("stork")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c2)
}
