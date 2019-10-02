package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gonuts/commander"
)

func makeCmdAdd(filename string) *commander.Command {
	cmdAdd := func(cmd *commander.Command, args []string) error {
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}
		w, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		defer w.Close()
		task := strings.Join(args, " ")
		_, err = fmt.Fprintln(w, task)
		fmt.Printf("Task added: %s\n", task)
		return err
	}

	return &commander.Command{
		Run:       cmdAdd,
		UsageLine: "add [message]",
		Short:     "add new todo",
	}
}
