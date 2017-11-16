package main

import (
	"fmt"
	"os"

	"github.com/gonuts/commander"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dir_db  = "/.todo"
	name_db = "/todo.db"
)

func main() {

	var (
		path_todo_db string
		err          error
	)

	path_todo_db = os.Getenv("HOME")

	path_todo_db += dir_db

	// path/to/whatever does not exist
	if _, err := os.Stat(path_todo_db); os.IsNotExist(err) {

		os.Mkdir(path_todo_db, 0755)

	}

	command := &commander.Command{

		UsageLine: os.Args[0],

		Short: "todo for cli",
	}

	command.Subcommands = []*commander.Command{
		make_cmd_list(path_todo_db + name_db),
		make_cmd_status(path_todo_db + name_db),
		make_cmd_add(path_todo_db + name_db),
		make_cmd_delete(path_todo_db + name_db),
		make_cmd_done(path_todo_db + name_db),
		make_cmd_undone(path_todo_db + name_db),
		make_cmd_clean(path_todo_db + name_db),
	}
	err = command.Dispatch(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	os.Exit(1)
}
