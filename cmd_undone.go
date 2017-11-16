package main

import (
	"github.com/gonuts/commander"
)

func make_cmd_undone(filename string) *commander.Command {
	cmd_undone := func(cmd *commander.Command, args []string) error {
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}

		db, err := InitDB(filename)
		defer db.Close()

		if err != nil {

			return err

		}
		if len(args) > 2 {
			for i := 1; i < len(args); i++ {
				err = UnCheckTodo(db, args[0], args[i])
				if err != nil {
					return err
				}

			}
		} else if len(args) > 1 {

			return UnCheckTodo(db, args[0], args[1])

		}
		return nil

	}

	return &commander.Command{
		Run:       cmd_undone,
		UsageLine: "undone [ID]",
		Short:     "undone the todo",
	}
}
