package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gonuts/commander"
)

func makeCmdUpdate(filename string) *commander.Command {
	cmdUpdate := func(cmd *commander.Command, args []string) error {
		if len(args) < 2 {
			cmd.Usage()
			return nil
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		task := strings.Join(args[1:], " ")

		w, err := os.Create(filename + "_")
		if err != nil {
			return err
		}
		defer w.Close()
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		br := bufio.NewReader(f)
		for n := 1; ; n++ {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			match := false
			if id == n {
				match = true
			}

			originalTask := string(b)
			if match {
				hasCompleted := strings.HasPrefix(originalTask, "-")
				if hasCompleted {
					task = "-" + task
				}
				_, err = fmt.Fprintf(w, "%s\n", task)
				if err != nil {
					return err
				}
				fmt.Printf("Task %d updated with message: %s\n", id, task)
			} else {
				_, err = fmt.Fprintf(w, "%s\n", originalTask)
				if err != nil {
					return err
				}
			}
		}
		f.Close()
		w.Close()
		err = os.Remove(filename)
		if err != nil {
			return err
		}
		return os.Rename(filename+"_", filename)
	}

	return &commander.Command{
		Run:       cmdUpdate,
		UsageLine: "update [ID] [message]",
		Short:     "update todo",
	}
}
