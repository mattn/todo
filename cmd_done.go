package main

import (
	"bufio"
	"fmt"
	"github.com/gonuts/commander"
	"io"
	"os"
	"strconv"
	"strings"
)

func make_cmd_done(filename string) *commander.Command {
	cmd_done := func(cmd *commander.Command, args []string) error {
		ids := []int{}
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
			line := strings.TrimSpace(string(b))
			if match && !strings.HasPrefix(line, "-") {
				_, err = fmt.Fprintf(w, "-%s\n", line)
				if err != nil {
					return err
				}
			} else {
				_, err = fmt.Fprintf(w, "%s\n", line)
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
		Run:       cmd_done,
		UsageLine: "done [ID]",
		Short:     "done the todo",
	}
}
