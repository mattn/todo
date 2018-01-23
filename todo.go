package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"path/filepath"
)

const (
	todo_filename = ".todo"
)

func main() {
	filename := ""
	existCurTodo := false
	curDir, err := os.Getwd()
	if err == nil {
		filename = filepath.Join(curDir, todo_filename)
		_, err = os.Stat(filename)
		if err == nil {
			existCurTodo = true
		}
	}
	if !existCurTodo {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		filename = filepath.Join(home, todo_filename)
	}
	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "todo for cli",
	}
	command.Subcommands = []*commander.Command{
		make_cmd_list(filename),
		make_cmd_add(filename),
		make_cmd_delete(filename),
		make_cmd_done(filename),
		make_cmd_undone(filename),
		make_cmd_clean(filename),
		make_cmd_sort(filename),
	}
	err = command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
