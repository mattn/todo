package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

const (
	time_format = "2006-01-02 15:04:05"
)

func make_cmd_status(filename string) *commander.Command {

	cmd_status := func(cmd *commander.Command, args []string) error {

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

				red := color.New(color.FgRed).Add(color.Bold).SprintFunc()

				black := color.New(color.FgBlack).Add(color.Bold).Add(color.Underline).SprintFunc()

				t, err := time.Parse(time_format, v.Status)

				if err != nil {

					return err

				}

				fmt.Printf("%03d: %s %s %s %s %s\n", v.Id, red("Task"), black(v.Todo), red("started"), black(Round(time.Since(t), time.Second)), red("ago"))

			case 1:

				blue := color.New(color.FgBlue).Add(color.Bold).SprintFunc()
				fmt.Printf("%03d:%s %s %s\n", v.Id, blue("Task"), v.Todo, blue("completed!"))

			}
		}

		return nil

	}

	flg := *flag.NewFlagSet("status", flag.ExitOnError)

	flg.Bool("n", false, "only not done")

	return &commander.Command{
		Run:       cmd_status,
		UsageLine: "status [options]",
		Short:     "show status of todos",
	}
}

func Round(d, r time.Duration) time.Duration {
	if r <= 0 {
		return d
	}
	neg := d < 0
	if neg {
		d = -d
	}
	if m := d % r; m+m < r {
		d = d - m
	} else {
		d = d + r - m
	}
	if neg {
		return -d
	}
	return d
}
