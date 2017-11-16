package main

import (
	"github.com/gonuts/commander"
)

func make_cmd_done(filename string) *commander.Command {
	cmd_done := func(cmd *commander.Command, args []string) error {
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
				err = CheckTodo(db, args[0], args[i])
				if err != nil {
					return err
				}

			}
		} else if len(args) > 1 {

			return CheckTodo(db, args[0], args[1])

		}
		return nil
	}

	return &commander.Command{
		Run:       cmd_done,
		UsageLine: "done [ID]",
		Short:     "done the todo",
	}
}
