package internal

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lukaszzieba/go-blog-agregator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, c Command) error {
		user, err := s.db.GetUser(context.Background(), s.Config.Current_user.Name)
		fmt.Println(user)
		if err != nil {
			return err
		}

		fmt.Println(user)
		if user.ID == uuid.Nil {
			return fmt.Errorf("user not found")
		}

		err = handler(s, c, user)
		if err != nil {
			return err
		}

		return nil
	}
}
