package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gonuts/commander"
)

func makeCmdDelete(filename string) *commander.Command {
	cmdDelete := func(cmd *commander.Command, args []string) error {
		if len(args) == 0 {
			cmd.Usage()
			return nil
		}
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
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
		n := 1
		for {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			match := false
			for _, id := range ids {
				if id == n {
					match = true
				}
			}
			if !match {
				_, err = fmt.Fprintf(w, "%s\n", string(b))
				if err != nil {
					return err
				}
			}
			n++
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
		Run:       cmdDelete,
		UsageLine: "delete [ID]",
		Short:     "delete the todo",
	}
}
