package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

const (
	done_mark1 = "\u2610"
	done_mark2 = "\u2611"
)

func make_cmd_list(filename string) *commander.Command {

	cmd_list := func(cmd *commander.Command, args []string) error {

		db, err := InitDB(filename)

		if err != nil {
			return err
		}
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}

		todos, err2 := ReadTodos(db, args[0])
		if err2 != nil {
			return err
		}
		for _, v := range todos {

			switch v.Done {
			case 0:
				fmt.Printf("%s %03d: %s  %s \n", done_mark1, v.Id, v.Todo, v.Status)

			case 1:
				fmt.Printf("%s %03d: %s\n", done_mark2, v.Id, v.Todo)

			}
		}

		return nil

	}

	flg := *flag.NewFlagSet("list", flag.ExitOnError)

	flg.Bool("n", false, "only not done")

	return &commander.Command{
		Run:       cmd_list,
		UsageLine: "list [options]",
		Short:     "show list index",
		Flag:      flg,
	}
}
