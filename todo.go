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
	todoDir := getTodoDir()
	filename := filepath.Join(todoDir, todo_filename)
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
	}
	err := command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func getTodoDir() string {
	dir, err := os.Getwd()
	if err == nil {
		filename := filepath.Join(dir, todo_filename)
		_, err = os.Stat(filename)
		if err == nil {
			return dir
		}
	}

	dir = os.Getenv("TODO_DIR")
	if dir != "" {
		return dir
	}

	dir = os.Getenv("HOME")
	if dir != "" {
		return dir
	}

	dir = os.Getenv("USERPROFILE")
	return dir
}
