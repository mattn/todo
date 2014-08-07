package main

import (
	"fmt"
	"github.com/gonuts/commander"
	"os"
	"path/filepath"
)

func main() {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE")
	}
	filename := filepath.Join(home, ".todo")

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
	}
	err := command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
