package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gonuts/commander"
)

const (
	todoFilename = ".todo"
)

func main() {
	filename := ""
	existCurTodo := false
	curDir, err := os.Getwd()
	if err == nil {
		filename = filepath.Join(curDir, todoFilename)
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
		filename = filepath.Join(home, todoFilename)
	}
	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "todo for cli",
	}
	command.Subcommands = []*commander.Command{
		makeCmdList(filename),
		makeCmdAdd(filename),
		makeCmdUpdate(filename),
		makeCmdDelete(filename),
		makeCmdDone(filename),
		makeCmdUndone(filename),
		makeCmdClean(filename),
		makeCmdSort(filename),
	}
	err = command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
