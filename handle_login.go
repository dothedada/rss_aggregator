package main

import "fmt"

func handlerLogin(s *State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login command expects a single parameter")
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Printf("Hello %s, welcome to your feed\n", cmd.args[0])
	return nil
}
