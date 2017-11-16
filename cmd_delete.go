package main

import (
	"github.com/gonuts/commander"
)

func make_cmd_delete(filename string) *commander.Command {
	cmd_delete := func(cmd *commander.Command, args []string) error {
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}

		db, err := InitDB(filename)
		defer db.Close()

		if err != nil {

			return err

		}
		if len(args) == 1 {

			return DeleteProject(db, args[0])
		} else {
			return DeleteTodo(db, args[0], args[1])
		}
		return nil

	}

	return &commander.Command{
		Run:       cmd_delete,
		UsageLine: "delete [ID]",
		Short:     "delete the todo",
	}
}
