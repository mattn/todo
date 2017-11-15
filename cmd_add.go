package main

import (
	"strings"

	"github.com/gonuts/commander"
)

func make_cmd_add(filename string) *commander.Command {
	cmd_add := func(cmd *commander.Command, args []string) error {
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}

		db, err := InitDB(filename)
		defer db.Close()

		if err != nil {

			return err

		}

		err = CreateTable(db, args[0])

		if err != nil {

			return err

		}

		if len(args) > 1 {
			return StoreTodo(db, args[0], strings.Join(args[1:], " "))
		}

		return nil

	}

	return &commander.Command{
		Run:       cmd_add,
		UsageLine: "add [message]",
		Short:     "add new todo",
	}
}
