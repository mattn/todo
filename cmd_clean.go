package main

import (
	"github.com/gonuts/commander"
)

func make_cmd_clean(filename string) *commander.Command {
	cmd_clean := func(cmd *commander.Command, args []string) error {
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}

		db, err := InitDB(filename)
		defer db.Close()

		if err != nil {

			return err

		}
		return DeleteAllTodos(db, args[0])

	}

	return &commander.Command{
		Run:       cmd_clean,
		UsageLine: "clean",
		Short:     "remove all done items",
	}
}
