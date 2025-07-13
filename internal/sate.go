package internal

import "fmt"

type state struct {
	Config *Config
}

func NewState(c *Config) *state {
	return &state{Config: c}
}

func HandlerLogin(s *state, cmd Command) error {
	fmt.Println(cmd.Args)
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login cmd requires some args")
	}

	userName := cmd.Args[0]
	_, err := s.Config.SetUser(userName)
	if err != nil {
		return err
	}

	return nil
}
